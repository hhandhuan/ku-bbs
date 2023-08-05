package model

import (
	"github.com/hhandhuan/ku-bbs/pkg/mysql"
	"gorm.io/gorm"
)

type IntegralLogs struct {
	Model
	UserId   uint64 `gorm:"column:user_id" db:"user_id" json:"user_id" form:"user_id"`         //用户 ID
	TargetId uint64 `gorm:"column:target_id" db:"target_id" json:"target_id" form:"target_id"` //目标 ID
	Rewards  int64  `gorm:"column:rewards" db:"rewards" json:"rewards" form:"rewards"`         //奖励积分
	Mode     string `gorm:"column:mode" db:"mode" json:"mode" form:"mode"`                     //获取方式
}

func IntegralLog() *gorm.DB {
	return mysql.GetInstance().Model(&IntegralLogs{})
}
