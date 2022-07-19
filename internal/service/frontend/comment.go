package frontend

import (
	"errors"
	"github.com/hhandhuan/ku-bbs/internal/consts"
	"github.com/hhandhuan/ku-bbs/internal/entity/frontend"
	remindSub "github.com/hhandhuan/ku-bbs/internal/subject/remind"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/hhandhuan/ku-bbs/internal/model"
	"github.com/hhandhuan/ku-bbs/internal/service"
)

func CommentService(ctx *gin.Context) *sComment {
	return &sComment{ctx: service.Context(ctx)}
}

type sComment struct {
	ctx *service.BaseContext
}

// Submit 提交评论
func (s *sComment) Submit(req *frontend.SubmitCommentReq) (uint64, error) {
	comment := &model.Comments{
		TopicId:   req.TopicId,
		ReplyId:   req.ReplyId,
		TargetId:  req.TargetId,
		UserId:    s.ctx.Auth().ID,
		Content:   req.Content,
		MDContent: req.MDContent,
	}
	r := model.Comment().M.Create(comment)
	if r.Error != nil {
		return 0, errors.New("服务内部错误")
	}
	if r.RowsAffected <= 0 {
		return 0, errors.New("提交失败，请稍后在试")
	}

	data := map[string]interface{}{
		"reply_id":      s.ctx.Auth().ID,
		"comment_count": gorm.Expr("comment_count + ?", 1),
		"last_reply_at": time.Now(),
	}

	r = model.Topic().M.Where("id = ?", req.TopicId).Updates(data)
	if r.Error != nil {
		return 0, errors.New("服务内部错误")
	}
	if r.RowsAffected <= 0 {
		return 0, errors.New("提交失败，请稍后在试")
	}

	if req.ReplyId <= 0 {
		sub := remindSub.New()
		sub.Attach(&remindSub.CommentObs{
			TopicID:   req.TopicId,
			Sender:    s.ctx.Auth().ID,
			CommentId: comment.ID,
		})
		sub.Notify()
	} else {
		sub := remindSub.New()
		sub.Attach(&remindSub.ReplyObs{
			TopicID:   req.TopicId,
			Sender:    s.ctx.Auth().ID,
			CommentId: comment.ID,
			Receiver:  req.ReplyId,
		})
		sub.Notify()
	}

	return comment.ID, nil
}

// GetList 获取列表
func (s *sComment) GetList(topicId uint64) ([]*frontend.Comment, error) {
	var list []*frontend.Comment

	query := model.Comment().M
	if s.ctx.Check() {
		query = query.Preload("Like", "user_id = ? AND source_type = ?", s.ctx.Auth().ID, consts.CommentSource)
	}

	r := query.
		Where("topic_id = ?", topicId).
		Order("id ASC").
		Preload("Publisher").
		Preload("Topic").
		Preload("Responder").
		Find(&list)
	if r.Error != nil {
		return nil, r.Error
	}

	return list, nil
}
