package backend

import "github.com/huhaophp/hblog/internal/model"

type GetNoticeListReq struct {
	Page     int    `form:"page"`
	Keywords string `form:"keywords"`
}

type SystemNotices struct {
	model.SystemNotices
	Publisher  model.Users  `gorm:"foreignKey:user_id"`
	TargetUser *model.Users `gorm:"foreignKey:target_id"`
}

type PublishNoticeReq struct {
	Title     string `v:"required#请输入消息标题" form:"title"`
	TargetId  string `form:"target_id"`
	Content   string `v:"required#请输入消息内容" form:"content"`
	MDContent string `v:"required#请输入消息内容" form:"md_content"`
}
