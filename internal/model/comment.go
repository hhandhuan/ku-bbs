package model

import (
	"github.com/hhandhuan/ku-bbs/pkg/db"
	"gorm.io/gorm"
)

type Comments struct {
	Model
	UserId    uint64 `gorm:"column:user_id" db:"user_id" json:"user_id" form:"user_id"`             //评论人 ID
	ReplyId   uint64 `gorm:"column:reply_id" db:"reply_id" json:"reply_id" form:"reply_id"`         //回复人 ID
	TopicId   uint64 `gorm:"column:topic_id" db:"topic_id" json:"topic_id" form:"topic_id"`         //话题 ID
	TargetId  uint64 `gorm:"column:target_id" db:"target_id" json:"target_id" form:"target_id"`     //回复目标 ID
	Content   string `gorm:"column:content" db:"content" json:"content" form:"content"`             //回复内容
	MDContent string `gorm:"column:md_content" db:"md_content" json:"md_content" form:"md_content"` //MD内容
	LikeCount uint64 `gorm:"column:like_count" db:"like_count" json:"like_count" form:"like_count"` //喜欢统计
}

type commentModel struct {
	M     *gorm.DB
	Table string
}

func Comment() *commentModel {
	return &commentModel{M: db.DB.Model(&Comments{}), Table: "comments"}
}
