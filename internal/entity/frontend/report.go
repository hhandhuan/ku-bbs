package frontend

type SubmitReportReq struct {
	Remark     string `v:"required|length:1,150#请输入举报内容|内容长度不能超过150个字符" form:"remark"`
	SourceID   uint64 `v:"required#请输入资源ID" form:"source_id"`
	SourceType string `v:"required|in:comment,topic#资源类型错误|资源类型错误" form:"source_type"`
	TargetID   uint64 `v:"required#请输入资源ID" form:"target_id"`
}
