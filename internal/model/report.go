package model

import (
	"github.com/hhandhuan/ku-bbs/pkg/mysql"
	"gorm.io/gorm"
)

type Reports struct {
	Model
	UserId     uint64 `gorm:"column:user_id" db:"user_id" json:"user_id" form:"user_id"`                 //举报人
	Remark     string `gorm:"column:remark" db:"remark" json:"remark" form:"remark"`                     //举报备注
	TargetId   uint64 `gorm:"column:target_id" db:"target_id" json:"target_id" form:"target_id"`         //被举报人ID
	SourceId   uint64 `gorm:"column:source_id" db:"source_id" json:"source_id" form:"source_id"`         //目标ID
	SourceType string `gorm:"column:source_type" db:"source_type" json:"source_type" form:"source_type"` //目标类型
	SourceUrl  string `gorm:"column:source_url" db:"source_url" json:"source_url" form:"source_url"`     //目标链接
	HandlerId  uint64 `gorm:"column:handler_id" db:"handler_id" json:"handler_id" form:"handler_id"`     //处理人ID
	State      uint8  `gorm:"column:state" db:"state" json:"state" form:"state"`                         //状态:0-待处理/1-已处理
}

func Report() *gorm.DB {
	return mysql.GetInstance().Model(&Reports{})
}
