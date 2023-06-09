package model

import (
	"github.com/hhandhuan/ku-bbs/pkg/mysql"
	"gorm.io/gorm"
)

type Follows struct {
	NoDeleteModel
	UserId   uint64 `gorm:"column:user_id" db:"user_id" json:"user_id" form:"user_id"`         //用户 ID
	TargetId uint64 `gorm:"column:target_id" db:"target_id" json:"target_id" form:"target_id"` //被关注用户 ID
	State    int8   `gorm:"column:state" db:"state" json:"state" form:"state"`                 //状态:1-关注/0-取消
}

type FollowModel struct {
	M     *gorm.DB
	table string
}

func Follow() *FollowModel {
	return &FollowModel{M: mysql.GetInstance().Model(&Follows{}), table: "follows"}
}
