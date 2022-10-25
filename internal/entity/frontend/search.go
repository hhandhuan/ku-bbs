package frontend

type GetSearchListReq struct {
	Page     int    `form:"page"`
	Keywords string `v:"required#搜索关键词不能为空" form:"keywords"`
}
