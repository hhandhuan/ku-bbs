package model

import (
	"github.com/hhandhuan/ku-bbs/pkg/mysql"
	"gorm.io/gorm"
	"time"
)

type Checkins struct {
	NoDeleteModel
	UserId         uint64    `gorm:"column:user_id" db:"user_id" json:"user_id" form:"user_id"`                                 //用户 ID
	CumulativeDays uint64    `gorm:"column:cumulative_days" db:"cumulative_days" json:"cumulative_days" form:"cumulative_days"` //累计签到(天)
	ContinuityDays uint64    `gorm:"column:continuity_days" db:"continuity_days" json:"continuity_days" form:"continuity_days"` //连续签到(天)
	LastTime       time.Time `gorm:"column:last_time" db:"last_time" json:"last_time" form:"last_time"`                         //最后签到时间
}

func Checkin() *gorm.DB {
	return mysql.GetInstance().Model(&Checkins{})
}
