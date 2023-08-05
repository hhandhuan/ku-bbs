package remind

import (
	"fmt"
	"log"

	"github.com/hhandhuan/ku-bbs/internal/consts"
	"github.com/hhandhuan/ku-bbs/internal/model"
)

// LikeObs 点赞提醒
type LikeObs struct {
	SourceID   uint64
	SourceType string
	Sender     uint64
	Receiver   uint64
}

// Update 点赞提醒
func (o *LikeObs) Update() {
	// 用户自身评论不写提醒消息
	if o.Sender == o.Receiver {
		return
	}

	var (
		sourceUrl     = ""
		sourceContent = ""
		action        = ""

		topic   *model.Topics
		comment *model.Comments
	)

	if o.SourceType == "comment" {
		if f := model.Comment().Where("id", o.SourceID).Find(&comment); f.Error != nil || comment == nil {
			log.Println(f.Error)
			return
		}
		if f := model.Topic().Where("id", comment.TopicId).Find(&topic); f.Error != nil || topic == nil {
			log.Println(f.Error)
			return
		}

		action = consts.LikeCommentRemind
		sourceContent = topic.Title
		sourceUrl = fmt.Sprintf("/topics/%d?j=comment%d", comment.TopicId, comment.ID)
	} else {
		f := model.Topic().Where("id", o.SourceID).Find(&topic)
		if f.Error != nil || topic == nil {
			log.Println(f.Error)
			return
		}

		action = consts.LikeTopicRemind
		sourceContent = topic.Title
		sourceUrl = fmt.Sprintf("/topics/%d", topic.ID)
	}

	c := model.Remind().Create(&model.Reminds{
		Sender:        o.Sender,
		Receiver:      o.Receiver,
		SourceId:      o.SourceID,
		SourceContent: sourceContent,
		SourceType:    o.SourceType,
		SourceUrl:     sourceUrl,
		Action:        action,
	})

	if c.Error != nil {
		log.Println(c.Error)
	}
}
