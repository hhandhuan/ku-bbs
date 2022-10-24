package frontend

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/frame/g"
	fe "github.com/hhandhuan/ku-bbs/internal/entity/frontend"
	"github.com/hhandhuan/ku-bbs/internal/service"
	"github.com/hhandhuan/ku-bbs/internal/service/frontend"
)

var User = cUser{}

type cUser struct{}

// HomePage 用户主页
func (c *cUser) HomePage(ctx *gin.Context) {
	s := service.Context(ctx)

	var req fe.GetUserHomeReq
	if err := ctx.ShouldBind(&req); err != nil {
		s.To("/").WithError(err).Redirect()
		return
	}
	if data, err := frontend.UserService(ctx).Home(&req); err != nil {
		s.To("/").WithError(err).Redirect()
	} else {
		s.View("frontend.user.home", data)
	}
}

// EditPage 用户编辑
func (c *cUser) EditPage(ctx *gin.Context) {
	s := service.Context(ctx)
	t := ctx.DefaultQuery("tab", "info")

	if !s.Check() {
		s.To("/login").WithError("请登录后，再继续操作").Redirect()
	} else {
		s.View("frontend.user.edit", gin.H{"tab": t})
	}
}

// EditSubmit 用户编辑
func (c *cUser) EditSubmit(ctx *gin.Context) {
	s := service.Context(ctx)
	t := ctx.DefaultQuery("tab", "info")
	p := ctx.Request.RequestURI

	switch t {
	case "info":
		var req fe.EditUserReq
		if err := ctx.ShouldBind(&req); err != nil {
			s.Back().WithError(err).Redirect()
			return
		}

		if err := g.Validator().Data(req).Run(context.Background()); err != nil {
			s.To(p).WithError(err.FirstError()).Redirect()
			return
		}

		if err := frontend.UserService(ctx).Edit(&req); err != nil {
			s.To(p).WithError(err).Redirect()
		} else {
			s.Back().WithMsg("修改信息成功").Redirect()
		}
	case "pass":
		var req fe.EditPasswordReq
		if err := ctx.ShouldBind(&req); err != nil {
			s.Back().WithError(err).Redirect()
			return
		}

		if err := g.Validator().Data(req).Run(context.Background()); err != nil {
			s.To(p).WithError(err.FirstError()).Redirect()
			return
		}

		if err := frontend.UserService(ctx).EditPassword(&req); err != nil {
			s.To(p).WithError(err).Redirect()
		} else {
			s.To("/login").WithMsg("修改密码成功，请重新登录").Redirect()
		}
	case "avatar":
		if err := frontend.UserService(ctx).EditAvatar(ctx); err != nil {
			s.To(p).WithError(err).Redirect()
		} else {
			s.Back().WithMsg("修改头像成功").Redirect()
		}
	}
}
