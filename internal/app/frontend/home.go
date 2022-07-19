package frontend

import (
	"github.com/gin-gonic/gin"
	fe "github.com/huhaophp/hblog/internal/entity/frontend"
	sv "github.com/huhaophp/hblog/internal/service"
	"github.com/huhaophp/hblog/internal/service/frontend"
)

var Home = home{}

type home struct{}

func (c *home) HomePage(ctx *gin.Context) {
	s := sv.Context(ctx)

	var req fe.GetTopicListReq
	if err := ctx.ShouldBind(&req); err != nil {
		s.Back().WithError(err).Redirect()
		return
	}

	data, _ := frontend.TopicService(ctx).GetList(&req)
	nodes, _ := frontend.NodeService(ctx).GetEnableNodes()
	checked, _ := frontend.CheckinService(ctx).IsCheckin()

	data["nodes"] = nodes
	data["checked"] = checked

	s.View("frontend.home.index", data)
}
