package frontend

import "github.com/huhaophp/hblog/internal/model"

type SystemUserNotice struct {
	model.SystemUserNotices
	Notice model.SystemNotices `gorm:"foreignKey:notice_id"`
}
