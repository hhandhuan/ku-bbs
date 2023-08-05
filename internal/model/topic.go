package model

import (
	"github.com/hhandhuan/ku-bbs/internal/model/column"
	"github.com/hhandhuan/ku-bbs/pkg/mysql"
	"gorm.io/gorm"
	"time"
)

type Topics struct {
	Model
	NodeId       int64      `gorm:"column:node_id" db:"node_id" json:"node_id" form:"node_id"`                         //分类 ID
	UserId       uint64     `gorm:"column:user_id" db:"user_id" json:"user_id" form:"user_id"`                         //用户 ID
	ReplyId      uint64     `gorm:"column:reply_id" db:"reply_id" json:"reply_id" form:"reply_id"`                     //最后回复者ID
	Title        string     `gorm:"column:title" db:"title" json:"title" form:"title"`                                 //话题标题
	Tags         column.SA  `gorm:"column:tags" db:"tags" json:"tags" form:"tags"`                                     //话题标签
	Images       column.SA  `gorm:"column:images" db:"images" json:"images" form:"images"`                             //图片集合
	CommentCount uint64     `gorm:"column:comment_count" db:"comment_count" json:"comment_count" form:"comment_count"` //评论统计
	ViewCount    uint64     `gorm:"column:view_count" db:"view_count" json:"view_count" form:"view_count"`             //浏览统计
	LikeCount    uint64     `gorm:"column:like_count" db:"like_count" json:"like_count" form:"like_count"`             //喜欢统计
	State        uint8      `gorm:"column:state" db:"state" json:"state" form:"state"`                                 //话题状态: 0-暂存/1-发布
	CommentState uint8      `gorm:"column:comment_state" db:"comment_state" json:"comment_state" form:"comment_state"` //讨论状态:0-关闭/1-开启
	Type         uint8      `gorm:"column:type" db:"type" json:"type" form:"type"`                                     //话题类型:0-默认/1-精华/2-置顶
	Brief        string     `gorm:"column:brief" db:"brief" json:"brief" form:"brief"`                                 //话题简介
	Content      string     `gorm:"column:content" db:"content" json:"content" form:"content"`                         //话题内容
	MDContent    string     `gorm:"column:md_content" db:"md_content" json:"md_content" form:"md_content"`             //MD内容
	LastReplyAt  *time.Time `gorm:"column:last_reply_at" db:"last_reply_at" json:"last_reply_at" form:"last_reply_at"` //最后回复时间
}

func Topic() *gorm.DB {
	return mysql.GetInstance().Model(&Topics{})
}
