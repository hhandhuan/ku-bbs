package model

import (
	"github.com/hhandhuan/ku-bbs/pkg/mysql"
	"gorm.io/gorm"
	"time"
)

type SystemUserNotices struct {
	Model
	UserId   uint64     `gorm:"column:user_id" db:"user_id" json:"user_id" form:"user_id"`         //用户 ID
	NoticeId uint64     `gorm:"column:notice_id" db:"notice_id" json:"notice_id" form:"notice_id"` //通知 ID
	ReadedAt *time.Time `gorm:"column:readed_at" db:"readed_at" json:"readed_at" form:"readed_at"` //阅读时间
}

func SystemUserNotice() *gorm.DB {
	return mysql.GetInstance().Model(&SystemUserNotices{})
}
