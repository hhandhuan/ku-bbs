package frontend

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/hhandhuan/ku-bbs/pkg/logger"
	"github.com/hhandhuan/ku-bbs/pkg/mysql"
	"github.com/hhandhuan/ku-bbs/pkg/redis"
	"github.com/hhandhuan/ku-bbs/pkg/utils"

	"github.com/gin-gonic/gin"
	"github.com/hhandhuan/ku-bbs/internal/consts"
	"github.com/hhandhuan/ku-bbs/internal/model"
	"github.com/hhandhuan/ku-bbs/internal/service"
	"gorm.io/gorm"
)

func CheckinService(ctx *gin.Context) *SCheckin {
	return &SCheckin{ctx: service.Context(ctx)}
}

type SCheckin struct {
	ctx *service.BaseContext
}

// Store  签到保存
func (s *SCheckin) Store() error {
	uid := s.ctx.Auth().ID

	key := fmt.Sprintf("user:%d:checkin", uid)
	exist, err := redis.GetInstance().SetNX(context.Background(), key, 1, time.Second*10).Result()
	if err != nil {
		logger.GetInstance().Error().Msgf("set setnx error: %v", err)
		return errors.New("请求太频繁")
	}
	if !exist {
		return errors.New("请求太频繁")
	}

	defer redis.GetInstance().Del(context.Background(), key)

	var checkin model.Checkins
	err = model.Checkin().Where("user_id", uid).Find(&checkin).Error
	if err != nil {
		logger.GetInstance().Error().Msgf("find checkin error: %v", err)
		return err
	}

	if checkin.ID > 0 && utils.ToDateTimeString(checkin.LastTime) >= utils.ToDateString(time.Now()) {
		return errors.New("请勿重复签到")
	}

	err = mysql.GetInstance().Transaction(func(tx *gorm.DB) error {
		if checkin.ID > 0 {
			preDate := utils.ToDateString(time.Now().AddDate(0, 0, -1))
			data := map[string]interface{}{
				"cumulative_days": gorm.Expr("cumulative_days + 1"),
				"last_time":       time.Now(),
			}
			// 判断是否连续签到
			if utils.ToDateTimeString(checkin.LastTime) == preDate {
				data["continuity_days"] = gorm.Expr("continuity_days + 1")
			} else {
				data["continuity_days"] = 1
			}
			u := tx.Model(&model.Checkins{}).Where("id", checkin.ID).Where("last_time", checkin.LastTime).Updates(data)
			if u.Error != nil {
				logger.GetInstance().Error().Msgf("update checkins error: %v", err)
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
			if c.Error != nil {
				logger.GetInstance().Error().Msgf("update checkins error: %v", err)
				return errors.New("签到失败")
			}
		}

		// 记录积分奖励日志
		c := tx.Model(&model.IntegralLogs{}).Create(&model.IntegralLogs{
			UserId:  uid,
			Rewards: consts.CHECKINReward,
			Mode:    consts.CheckinMode,
		})
		if c.Error != nil {
			logger.GetInstance().Error().Msgf("record integral logs error: %v", err)
			return errors.New("签到失败")
		}

		// 更新用户积分数
		u := tx.Model(&model.Users{}).Where("id", uid).Updates(map[string]interface{}{
			"integral": gorm.Expr("integral + ?", consts.CHECKINReward),
		})
		if u.Error != nil {
			logger.GetInstance().Error().Msgf("update users integral error: %v", err)
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
func (s *SCheckin) IsCheckin() (bool, error) {
	if !s.ctx.Check() {
		return false, nil
	}

	date := utils.ToDateString(time.Now())

	startAt := date + " 00:00:00"
	endedAt := date + " 23:59:59"

	var checkin *model.Checkins
	f := model.Checkin().
		Where("last_time >= ?", startAt).
		Where("last_time <= ?", endedAt).
		Where("user_id", s.ctx.Auth().ID).
		Find(&checkin)
	if f.Error != nil {
		return false, f.Error
	}

	return checkin.ID > 0, nil
}
