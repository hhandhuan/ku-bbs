package backend

import (
	"github.com/gin-gonic/gin"

	be "github.com/huhaophp/hblog/internal/entity/backend"
	sv "github.com/huhaophp/hblog/internal/service"
	bs "github.com/huhaophp/hblog/internal/service/backend"
)

var User = cUser{}

type cUser struct{}

// IndexPage 用户主页
func (c *cUser) IndexPage(ctx *gin.Context) {
	s := sv.Context(ctx)

	var req be.GetUserListReq
	if err := ctx.ShouldBind(&req); err != nil {
		s.Back().WithError(err).Redirect()
		return
	}

	if data, err := bs.UserService(ctx).GetList(&req); err != nil {
		s.Back().WithError(err).Redirect()
	} else {
		s.View("backend.user.index", data)
	}
}
