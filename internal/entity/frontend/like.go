package frontend

type LikeReq struct {
	SourceID     uint64 `v:"required#请输入资源ID" form:"source_id"`
	SourceType   string `v:"required|in:comment,topic#资源类型错误|资源类型错误" form:"source_type"`
	TargetUserID uint64 `v:"required#目标用户错误" form:"target_user_id"`
}
