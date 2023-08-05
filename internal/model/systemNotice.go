package model

import (
	"github.com/hhandhuan/ku-bbs/pkg/mysql"
	"gorm.io/gorm"
)

type SystemNotices struct {
	Model
	UserId    uint64 `gorm:"column:user_id" db:"user_id" json:"user_id" form:"user_id"`             //发布人 ID
	Title     string `gorm:"column:title" db:"title" json:"title" form:"title"`                     //消息标题
	TargetId  string `gorm:"column:target_id" db:"target_id" json:"target_id" form:"target_id"`     //接受者 ID
	Content   string `gorm:"column:content" db:"content" json:"content" form:"content"`             //消息内容
	MDContent string `gorm:"column:md_content" db:"md_content" json:"md_content" form:"md_content"` //markdown 内容
}

func SystemNotice() *gorm.DB {
	return mysql.GetInstance().Model(&SystemNotices{})
}
