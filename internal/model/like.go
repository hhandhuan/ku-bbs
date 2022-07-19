package model

import (
	"github.com/huhaophp/hblog/pkg/db"
	"gorm.io/gorm"
)

type Likes struct {
	NoDeleteModel
	UserId       uint64 `gorm:"column:user_id" db:"user_id" json:"user_id" form:"user_id"`                             //用户 ID
	TargetUserId uint64 `gorm:"column:target_user_id" db:"target_user_id" json:"target_user_id" form:"target_user_id"` //目标用户ID
	SourceId     uint64 `gorm:"column:source_id" db:"source_id" json:"source_id" form:"source_id"`                     //资源 ID
	SourceType   string `gorm:"column:source_type" db:"source_type" json:"source_type" form:"source_type"`             //资源类型:topic/comment
	State        uint8  `gorm:"column:state" db:"state" json:"state" form:"state"`                                     //状态: 0-否/1-是
}

type LikeModel struct {
	M     *gorm.DB
	table string
}

func Like() *LikeModel {
	return &LikeModel{M: db.DB.Model(&Likes{}), table: "likes"}
}
