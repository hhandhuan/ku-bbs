package frontend

import "github.com/hhandhuan/ku-bbs/internal/model"

type GetTopicListReq struct {
	Page int    `form:"page"`
	Type string `form:"type"`
}

type PublishTopicReq struct {
	Title     string `v:"required#请输入话题标题" form:"title"`
	NodeId    int64  `v:"required|integer|min:1#请选择话题分类|请选择话题分类|请选择话题分类" form:"node_id"`
	Content   string `v:"required#请输入话题内容" form:"content"`
	MDContent string `v:"required#请输入话题内容" form:"md_content"`
	Tags      string `form:"tags"`
}

type Topic struct {
	model.Topics
	Publisher User        `gorm:"foreignKey:user_id"`
	Responder model.Users `gorm:"foreignKey:reply_id"`
	Node      model.Nodes `gorm:"foreignKey:node_id"`
	Like      *Like       `gorm:"foreignKey:source_id"`
	Likes     []*Like     `gorm:"foreignKey:source_id"`
	PostDays  int         `gorm:"-"`
}
