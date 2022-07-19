package frontend

import "github.com/huhaophp/hblog/internal/model"

// SubmitCommentReq  评论话题
type SubmitCommentReq struct {
	ReplyId   uint64 `form:"reply_id"`
	TargetId  uint64 `form:"target_id"`
	TopicId   uint64 `v:"required#帖子ID错误" form:"topic_id"`
	Content   string `v:"required#请输入评论内容" form:"content"`
	MDContent string `v:"required#请输入评论内容" form:"md_content"`
}

// Comment 评论
type Comment struct {
	model.Comments
	Publisher model.Users  `gorm:"foreignKey:user_id"`
	Responder *model.Users `gorm:"foreignKey:reply_id"`
	Topic     model.Topics `gorm:"foreignKey:topic_id"`
	Like      *model.Likes `gorm:"foreignKey:source_id"`
}
