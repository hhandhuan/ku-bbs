package model

import (
	"github.com/hhandhuan/ku-bbs/pkg/db"
	"gorm.io/gorm"
	"time"
)

type Reminds struct {
	Model
	Sender        uint64     `gorm:"column:sender" db:"sender" json:"sender" form:"sender"`                                 //发送人ID
	Receiver      uint64     `gorm:"column:receiver" db:"receiver" json:"receiver" form:"receiver"`                         //接受者ID
	SourceId      uint64     `gorm:"column:source_id" db:"source_id" json:"source_id" form:"source_id"`                     //资源 ID
	SourceType    string     `gorm:"column:source_type" db:"source_type" json:"source_type" form:"source_type"`             //资源类型
	SourceContent string     `gorm:"column:source_content" db:"source_content" json:"source_content" form:"source_content"` //资源内容
	SourceUrl     string     `gorm:"column:source_url" db:"source_url" json:"source_url" form:"source_url"`                 //提醒发生地址
	Action        string     `gorm:"column:action" db:"action" json:"action" form:"action"`                                 //动作类型
	ReadedAt      *time.Time `gorm:"column:readed_at" db:"readed_at" json:"readed_at" form:"readed_at"`                     //阅读时间
}

type RemindModel struct {
	M     *gorm.DB
	Table string
}

func Remind() *RemindModel {
	return &RemindModel{M: db.DB.Model(&Reminds{}), Table: "reminds"}
}
