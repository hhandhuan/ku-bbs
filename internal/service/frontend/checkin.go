package frontend

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/hhandhuan/ku-bbs/internal/consts"
	"github.com/hhandhuan/ku-bbs/internal/model"
	"github.com/hhandhuan/ku-bbs/internal/service"
	"github.com/hhandhuan/ku-bbs/pkg/db"
	"gorm.io/gorm"
	"log"
	"time"
)

func CheckinService(ctx *gin.Context) *sCheckin {
	return &sCheckin{ctx: service.Context(ctx)}
}

type sCheckin struct {
	ctx *service.BaseContext
}

// Store  签到保存
func (s *sCheckin) Store() error {
	uid := s.ctx.Auth().ID

	var checkin model.Checkins
	f := model.Checkin().M.Where("user_id", uid).Find(&checkin)
	if f.Error != nil {
		return f.Error
	}

	if checkin.ID > 0 && checkin.LastTime.Format("2006-01-02") >= time.Now().Format("2006-01-02") {
		return errors.New("请勿重复签到")
	}

	err := db.DB.Transaction(func(tx *gorm.DB) error {
		if checkin.ID > 0 {
			preDate := time.Now().AddDate(0, 0, -1).Format("2006-01-02")
			data := map[string]interface{}{
				"cumulative_days": gorm.Expr("cumulative_days + 1"),
				"last_time":       time.Now(),
			}
			// 判断是否连续签到
			if checkin.LastTime.Format("2006-01-02") == preDate {
				data["continuity_days"] = gorm.Expr("continuity_days + 1")
			} else {
				data["continuity_days"] = 1
			}
			u := tx.Model(&model.Checkins{}).Where("id", checkin.ID).Where("last_time", checkin.LastTime).Updates(data)
			if u.Error != nil || u.RowsAffected <= 0 {
				log.Println(u.Error)
				return errors.New("签到失败")
			}
		} else {
			// 生成签到关联
			c := tx.Model(&model.Checkins{}).Create(&model.Checkins{
				UserId:         uid,
				CumulativeDays: 1,
				ContinuityDays: 1,
				LastTime:       time.Now(),
			})
			if c.Error != nil || c.RowsAffected <= 0 {
				log.Println(c.Error)
				return errors.New("签到失败")
			}
		}

		// 记录积分奖励日志
		c := tx.Model(&model.IntegralLogs{}).Create(&model.IntegralLogs{
			UserId:  uid,
			Rewards: consts.CHECKINReward,
			Mode:    consts.CheckinMode,
		})
		if c.Error != nil || c.RowsAffected <= 0 {
			log.Println(c.Error)
			return errors.New("签到失败")
		}

		// 更新用户积分数
		u := tx.Model(&model.Users{}).Where("id", uid).Updates(map[string]interface{}{
			"integral": gorm.Expr("integral + ?", consts.CHECKINReward),
		})
		if u.Error != nil || u.RowsAffected <= 0 {
			log.Println(u.Error)
			return errors.New("签到失败")
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

// IsCheckin  今日是否签到
func (s *sCheckin) IsCheckin() (bool, error) {
	if !s.ctx.Check() {
		return false, nil
	}

	date := time.Now().Format("2006-01-02")

	startAt := date + " 00:00:00"
	endedAt := date + " 23:59:59"

	var checkin *model.Checkins
	f := model.Checkin().M.
		Where("last_time >= ?", startAt).
		Where("last_time <= ?", endedAt).
		Where("user_id", s.ctx.Auth().ID).
		Find(&checkin)
	if f.Error != nil {
		return false, f.Error
	}

	return checkin.ID > 0, nil
}
