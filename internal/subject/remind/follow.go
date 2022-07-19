package remind

import (
	"github.com/huhaophp/hblog/internal/consts"
	"github.com/huhaophp/hblog/internal/model"
	"log"
)

// FollowObs 关注提醒
type FollowObs struct {
	Sender   uint64
	Receiver uint64
}

// Update 回复评论提醒
func (o *FollowObs) Update() {
	r := model.Remind().M.Create(&model.Reminds{
		Sender:        o.Sender,
		Receiver:      o.Receiver,
		SourceId:      0,
		SourceContent: "",
		SourceType:    consts.UserSource,
		SourceUrl:     "",
		Action:        consts.FollowUserRemind,
	})
	if r.Error != nil {
		log.Println(r.Error)
	}
}
