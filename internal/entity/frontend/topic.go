package frontend

import "github.com/hhandhuan/ku-bbs/internal/model"

type GetTopicListReq struct {
	Page int    `form:"page"`
	Type string `form:"type"`
}

type PublishTopicReq struct {
	Title   string `v:"required#请输入话题标题" form:"title"`
	NodeId  int64  `v:"required|integer#请选择话题分类|话题分类格式错误" form:"node_id"`
	Content string `v:"required#请输入话题内容" form:"content"`
}

type Topic struct {
	model.Topics
	Publisher User         `gorm:"foreignKey:user_id"`
	Responder model.Users  `gorm:"foreignKey:reply_id"`
	Node      model.Nodes  `gorm:"foreignKey:node_id"`
	Like      *model.Likes `gorm:"foreignKey:source_id"`
}
