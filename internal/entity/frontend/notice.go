package frontend

import "github.com/hhandhuan/ku-bbs/internal/model"

type SystemUserNotice struct {
	model.SystemUserNotices
	Notice model.SystemNotices `gorm:"foreignKey:notice_id"`
}
