package frontend

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/huhaophp/hblog/internal/consts"
	"github.com/huhaophp/hblog/internal/entity/frontend"
	"github.com/huhaophp/hblog/internal/model"
	"github.com/huhaophp/hblog/internal/service"
	remindSub "github.com/huhaophp/hblog/internal/subject/remind"
	"github.com/huhaophp/hblog/pkg/db"
	"gorm.io/gorm"
)

func LikeService(ctx *gin.Context) *sLike {
	return &sLike{ctx: service.Context(ctx)}
}

type sLike struct {
	ctx *service.BaseContext
}

// Like 点赞提交
func (s *sLike) Like(req *frontend.LikeReq) error {

	liked, err := s.IsLiked(req.SourceID, req.SourceType)
	if err != nil {
		return errors.New("点赞失败，请稍后在试")
	}

	if liked {
		return errors.New("无法重复点赞")
	}

	err = db.DB.Transaction(func(tx *gorm.DB) error {
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

		if req.SourceType == consts.TopicSource {
			u := tx.Model(&model.Topics{}).Where("id", req.SourceID).Updates(data)
			if u.Error != nil || u.RowsAffected <= 0 {
				return errors.New("点赞失败，请稍后在试")
			}
			return nil
		}

		u := tx.Model(&model.Comments{}).Where("id", req.SourceID).Updates(data)
		if u.Error != nil || u.RowsAffected <= 0 {
			return errors.New("点赞失败，请稍后在试")
		}

		return nil
	})

	if err != nil {
		return err
	}

	sub := remindSub.New()
	sub.Attach(&remindSub.LikeObs{
		Sender:     s.ctx.Auth().ID,
		Receiver:   req.TargetUserID,
		SourceID:   req.SourceID,
		SourceType: req.SourceType,
	})
	sub.Notify()

	return nil
}

// IsLiked 是否点赞
func (s *sLike) IsLiked(id uint64, source string) (bool, error) {
	user := s.ctx.Auth()

	var like *model.Likes
	f := model.Like().M.Where(&model.Likes{UserId: user.ID, SourceType: source, SourceId: id}).Find(&like)
	if f.Error != nil {
		return false, f.Error
	} else {
		return like.ID > 0, nil
	}
}
