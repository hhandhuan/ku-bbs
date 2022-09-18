package backend

type GetNodeListReq struct {
	Keywords string `form:"keywords"`
}

type CreateNodeReq struct {
	Name  string `v:"required#请输入节点名称" form:"title"`
	Alias string `v:"required#请输入节点别名" form:"alias"`
	Sort  uint8  `v:"required|integer#请填写排序|排序格式错误" form:"sort"`
	State uint8  `v:"required|in:0,1#请选择节点状态" form:"state"`
	Desc  string `v:"required#请输入节点简介" form:"desc"`
}

type EditNodeReq struct {
	Name  string `v:"required#请输入节点名称" form:"title"`
	Alias string `v:"required#请输入节点别名" form:"alias"`
	Sort  uint8  `v:"required|integer#请填写排序|排序格式错误" form:"sort"`
	State uint8  `v:"required|in:0,1#请选择节点状态" form:"state"`
	Desc  string `v:"required#请输入节点简介" form:"desc"`
}
