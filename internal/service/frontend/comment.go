package frontend

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/hhandhuan/ku-bbs/internal/consts"
	"github.com/hhandhuan/ku-bbs/internal/entity/frontend"
	remindSub "github.com/hhandhuan/ku-bbs/internal/subject/remind"
	"github.com/hhandhuan/ku-bbs/pkg/db"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/hhandhuan/ku-bbs/internal/model"
	"github.com/hhandhuan/ku-bbs/internal/service"
)

func CommentService(ctx *gin.Context) *SComment {
	return &SComment{ctx: service.Context(ctx)}
}

type SComment struct {
	ctx *service.BaseContext
}

// Submit 提交评论
func (s *SComment) Submit(req *frontend.SubmitCommentReq) (uint64, error) {

	var topic *model.Topics
	// 检查话题是否存在
	f := model.Topic().M.Where("id = ?", req.TopicId).Find(&topic)
	if f.Error != nil {
		log.Println("delete topic error: ", f.Error)
		return 0, f.Error
	}
	if topic.ID <= 0 {
		return 0, errors.New("话题资源未找到")
	}
	if topic.CommentState == consts.DisableState {
		return 0, errors.New("此话题已关闭讨论，不再接受任何回复")
	}

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
func (s *SComment) GetList(topicId, authorId uint64) ([]*frontend.Comment, error) {
	var list []*frontend.Comment

	query := model.Comment().M
	if s.ctx.Check() {
		query = query.Preload("Like", "user_id = ? AND source_type = ?", s.ctx.Auth().ID, consts.CommentSource)
	}

	if authorId > 0 {
		query = query.Where("user_id", authorId)
	}

	r := query.Unscoped().
		Where("topic_id = ?", topicId).
		Order("id ASC").
		Preload("Publisher").
		Preload("Topic").
		Preload("Responder").
		Find(&list)
	if r.Error != nil {
		return nil, r.Error
	}

	floorMap := make(map[uint64]int, len(list))
	for index, item := range list {
		floorMap[item.ID] = index + 1
		list[index].Floor = floorMap[item.ID]
		if item.TargetId > 0 {
			list[index].ReplyFloor = floorMap[item.TargetId]
		}
	}

	return list, nil
}

// Delete 删除评论
func (s *SComment) Delete(id uint64) error {
	if !s.ctx.Check() {
		return errors.New("权限不足")
	}

	var comment *model.Comments
	f := model.Comment().M.Where("id", id).Find(&comment)
	if f.Error != nil || comment == nil {
		log.Println(f.Error)
		return errors.New("删除失败")
	}

	if comment.UserId != s.ctx.Auth().ID {
		return errors.New("权限不足")
	}

	err := db.DB.Transaction(func(tx *gorm.DB) error {
		d := tx.Delete(&model.Comments{}, id)
		if d.Error != nil || d.RowsAffected <= 0 {
			return fmt.Errorf("delete comment error: %v", d.Error)
		}
		//u := tx.Model(&model.Topics{}).Where("id", comment.TopicId).Updates(map[string]interface{}{
		//	"comment_count": gorm.Expr("comment_count - 1"),
		//})
		//if u.Error != nil || u.RowsAffected <= 0 {
		//	return fmt.Errorf("delete comment error: %v", d.Error)
		//}
		return nil
	})

	return err
}
