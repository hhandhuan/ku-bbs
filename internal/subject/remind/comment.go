package remind

import (
	"errors"
	"fmt"
	"github.com/huhaophp/hblog/internal/consts"
	"github.com/huhaophp/hblog/internal/model"
	"gorm.io/gorm"
	"log"
)

// CommentObs 评论话题提醒
type CommentObs struct {
	Sender    uint64
	TopicID   uint64
	CommentId uint64
}

// Update 评论话题提醒
func (o *CommentObs) Update() {
	var topic model.Topics
	r := model.Topic().M.Where("id", o.TopicID).First(&topic)
	if r.Error != nil && !errors.Is(r.Error, gorm.ErrRecordNotFound) {
		log.Println(r.Error)
		return
	}

	// 用户自身评论不写提醒消息
	if o.Sender == topic.UserId {
		return
	}

	r = model.Remind().M.Create(&model.Reminds{
		Sender:        o.Sender,
		Receiver:      topic.UserId,
		SourceId:      topic.ID,
		SourceContent: topic.Title,
		SourceType:    model.Topic().Table,
		SourceUrl:     fmt.Sprintf("/topics/%d?j=comment%d", o.TopicID, o.CommentId),
		Action:        consts.CommentTopicRemind,
	})

	if r.Error != nil {
		log.Println(r.Error)
	}
}
