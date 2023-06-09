package frontend

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hhandhuan/ku-bbs/internal/consts"
	"github.com/hhandhuan/ku-bbs/internal/entity/frontend"
	"github.com/hhandhuan/ku-bbs/internal/model"
	"github.com/hhandhuan/ku-bbs/internal/service"
	remindSub "github.com/hhandhuan/ku-bbs/internal/subject/remind"
	"github.com/hhandhuan/ku-bbs/pkg/mysql"
	"github.com/hhandhuan/ku-bbs/pkg/redis"
	"gorm.io/gorm"
	"time"
)

func LikeService(ctx *gin.Context) *SLike {
	return &SLike{ctx: service.Context(ctx)}
}

type SLike struct {
	ctx *service.BaseContext
}

// Like 点赞提交
func (s *SLike) Like(req *frontend.LikeReq) error {

	auth := s.ctx.Auth()

	lockKey := fmt.Sprintf("user:%d:source:%d:like", auth.ID, req.SourceID)

	val, err := redis.GetInstance().SetNX(context.Background(), lockKey, 1, time.Second*10).Result()
	if err != nil {
		return errors.New("点赞失败，请稍后在试")
	}

	if !val {
		return errors.New("点赞失败, 操作太频繁")
	}

	defer redis.GetInstance().Del(context.Background(), lockKey)

	liked, err := s.IsLiked(req.SourceID, req.SourceType)
	if err != nil {
		return errors.New("点赞失败，请稍后在试")
	}

	if liked {
		return errors.New("无法重复点赞")
	}

	err = mysql.GetInstance().Transaction(func(tx *gorm.DB) error {
		c := tx.Create(&model.Likes{
			UserId:       s.ctx.Auth().ID,
			SourceType:   req.SourceType,
			SourceId:     req.SourceID,
			TargetUserId: req.TargetUserID,
			State:        consts.Liked,
		})
		if c.Error != nil || c.RowsAffected <= 0 {
			return errors.New("点赞失败，请稍后在试")
		}

		data := map[string]interface{}{
			"like_count": gorm.Expr("like_count + ?", 1),
		}

		switch req.SourceType {
		case consts.TopicSource:
			u := tx.Model(&model.Topics{}).Where("id", req.SourceID).Updates(data)
			if u.Error != nil || u.RowsAffected <= 0 {
				return errors.New("点赞失败，请稍后在试")
			}
		default:
			u := tx.Model(&model.Comments{}).Where("id", req.SourceID).Updates(data)
			if u.Error != nil || u.RowsAffected <= 0 {
				return errors.New("点赞失败，请稍后在试")
			}
		}

		return nil
	})

	if err != nil {
		return err
	}

	obs := &remindSub.LikeObs{
		Sender:     s.ctx.Auth().ID,
		Receiver:   req.TargetUserID,
		SourceID:   req.SourceID,
		SourceType: req.SourceType,
	}

	sub := remindSub.New()
	sub.Attach(obs)
	sub.Notify()

	return nil
}

// IsLiked 是否点赞
func (s *SLike) IsLiked(id uint64, source string) (bool, error) {
	user := s.ctx.Auth()

	var like *model.Likes
	f := model.Like().M.Where(&model.Likes{UserId: user.ID, SourceType: source, SourceId: id}).Find(&like)
	if f.Error != nil {
		return false, f.Error
	} else {
		return like.ID > 0, nil
	}
}
