package frontend

import (
	"context"
	fe "github.com/hhandhuan/ku-bbs/internal/entity/frontend"

	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/frame/g"

	"github.com/hhandhuan/ku-bbs/internal/service"
	"github.com/hhandhuan/ku-bbs/internal/service/frontend"
)

var Auth = auth{}

type auth struct{}

// RegisterPage 注册页面
func (c *auth) RegisterPage(ctx *gin.Context) {
	service.Context(ctx).View("frontend.auth.register", gin.H{})
}

// RegisterSubmit 注册提交
func (c *auth) RegisterSubmit(ctx *gin.Context) {
	s := service.Context(ctx)
	p := ctx.DefaultQuery("back", "/")

	var req fe.RegisterReq
	if err := ctx.ShouldBind(&req); err != nil {
		s.Back().WithError(err).Redirect()
		return
	}

	if err := g.Validator().Data(req).Run(context.Background()); err != nil {
		s.Back().WithError(err.FirstError()).Redirect()
		return
	}

	if err := frontend.UserService(ctx).Register(&req); err != nil {
		s.Back().WithError(err).Redirect()
	} else {
		s.To(p).WithMsg("注册成功， 欢迎来到酷社区").Redirect()
	}
}

// LoginPage 登录页面
func (c *auth) LoginPage(ctx *gin.Context) {
	p := ctx.DefaultQuery("back", "/")
	s := service.Context(ctx)

	if s.Check() {
		s.To(p).Redirect()
	} else {
		s.View("frontend.auth.login", gin.H{"path": p})
	}
}

// LoginSubmit 登录提交
func (c *auth) LoginSubmit(ctx *gin.Context) {
	s := service.Context(ctx)
	p := ctx.DefaultQuery("back", "/")

	var req fe.LoginReq
	if err := ctx.ShouldBind(&req); err != nil {
		s.Back().WithError(err).Redirect()
		return
	}

	if err := g.Validator().Data(req).Run(context.Background()); err != nil {
		s.Back().WithError(err.FirstError()).Redirect()
		return
	}

	if err := frontend.UserService(ctx).Login(&req); err != nil {
		s.Back().WithError(err).Redirect()
	} else {
		s.To(p).WithMsg("登录成功， 欢迎来到酷社区").Redirect()
	}
}

// LogoutSubmit 用户登出
func (c *auth) LogoutSubmit(ctx *gin.Context) {
	s := service.Context(ctx)

	frontend.UserService(ctx).Logout()

	s.To("/").WithMsg("退出成功").Redirect()
}
