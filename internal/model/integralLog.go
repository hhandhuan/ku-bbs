package model

import (
	"github.com/hhandhuan/ku-bbs/pkg/db"
	"gorm.io/gorm"
	"time"
)

type IntegralLogs struct {
	ID        uint64    `gorm:"primarykey"`                                                            // 主键ID
	UserId    uint64    `gorm:"column:user_id" db:"user_id" json:"user_id" form:"user_id"`             //用户 ID
	Rewards   int64     `gorm:"column:rewards" db:"rewards" json:"rewards" form:"rewards"`             //奖励积分
	Mode      string    `gorm:"column:mode" db:"mode" json:"mode" form:"mode"`                         //获取方式
	CreatedAt time.Time `gorm:"column:created_at" db:"created_at" json:"created_at" form:"created_at"` //创建时间
}

type integralLogModel struct {
	M     *gorm.DB
	Table string
}

func IntegralLog() *integralLogModel {
	return &integralLogModel{M: db.DB.Model(&IntegralLogs{}), Table: "integral_logs"}
}
