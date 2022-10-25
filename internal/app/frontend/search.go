package frontend

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/frame/g"
	fe "github.com/hhandhuan/ku-bbs/internal/entity/frontend"
	"github.com/hhandhuan/ku-bbs/internal/service"
	"github.com/hhandhuan/ku-bbs/internal/service/frontend"
)

var Search = cSearch{}

type cSearch struct{}

// List 搜索列表
func (c *cSearch) List(ctx *gin.Context) {
	s := service.Context(ctx)

	var req fe.GetSearchListReq
	if err := ctx.ShouldBind(&req); err != nil {
		s.To("/").WithError(err).Redirect()
		return
	}

	if err := g.Validator().Data(req).Run(context.Background()); err != nil {
		s.To("/").WithError(err.FirstError()).Redirect()
		return
	}

	data, err := frontend.SearchService(ctx).GetList(&req)
	if err != nil {
		s.To("/").WithError(err).Redirect()
	} else {
		s.WithData(req).View("frontend.search.list", data)
	}
}
