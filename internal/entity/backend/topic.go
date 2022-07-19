package backend

import "github.com/huhaophp/hblog/internal/model"

type GetTopicListReq struct {
	Page     int    `form:"page"`
	Keywords string `form:"keywords"`
	UserID   string `form:"user_id"`
}

type Topic struct {
	model.Topics
	Publisher model.Users `gorm:"foreignKey:user_id"`
	Node      model.Nodes `gorm:"foreignKey:node_id"`
}
