package frontend

import "github.com/hhandhuan/ku-bbs/internal/model"

// SubmitCommentReq  评论话题
type SubmitCommentReq struct {
	ReplyId   uint64 `form:"reply_id"`
	TargetId  uint64 `form:"target_id"`
	TopicId   uint64 `v:"required#帖子ID错误" form:"topic_id"`
	Content   string `v:"required#请输入评论内容" form:"content"`
	MDContent string `v:"required#请输入评论内容" form:"md_content"`
}

// DeleteCommentReq  删除评论
type DeleteCommentReq struct {
	ID uint64 `v:"required#参数错误" form:"id"`
}

// Comment 评论
type Comment struct {
	model.Comments
	Publisher model.Users  `gorm:"foreignKey:user_id"`   // 评论人
	Responder *model.Users `gorm:"foreignKey:reply_id"`  // 回复人
	Topic     model.Topics `gorm:"foreignKey:topic_id"`  // 话题
	Like      *model.Likes `gorm:"foreignKey:source_id"` // 点赞
}
