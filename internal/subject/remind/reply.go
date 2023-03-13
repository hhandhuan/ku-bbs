package remind

import (
	"fmt"
	"log"

	"github.com/hhandhuan/ku-bbs/internal/consts"
	"github.com/hhandhuan/ku-bbs/internal/model"
)

// ReplyObs 回复评论提醒
type ReplyObs struct {
	Sender    uint64
	Receiver  uint64
	TopicID   uint64
	CommentId uint64
}

// Update 回复评论提醒
func (o *ReplyObs) Update() {
	var topic *model.Topics
	r := model.Topic().M.Where("id", o.TopicID).Find(&topic)
	if r.Error != nil || topic.ID <= 0 {
		log.Println(r.Error)
		return
	}

	// 用户自身评论不写提醒消息
	if o.Sender == o.Receiver {
		return
	}

	r = model.Remind().M.Create(&model.Reminds{
		Sender:        o.Sender,
		Receiver:      o.Receiver,
		SourceId:      topic.ID,
		SourceContent: topic.Title,
		SourceType:    consts.TopicSource,
		SourceUrl:     fmt.Sprintf("/topics/%d?j=comment%d", o.TopicID, o.CommentId),
		Action:        consts.ReplyCommentRemind,
	})

	if r.Error != nil {
		log.Println(r.Error)
	}
}
