package route

import (
	"github.com/gin-gonic/gin"
	srv "github.com/hhandhuan/ku-bbs/internal/service"
)

func isAdmin(ctx *gin.Context) {
	s := srv.Context(ctx)
	if !s.IsAdmin() {
		s.To("/").WithError("无权限访问").Redirect()
		ctx.Abort()
		return
	} else {
		ctx.Next()
	}
}

func isAuth(ctx *gin.Context) {
	s := srv.Context(ctx)
	if !s.Check() {
		s.To("/login").WithError("请先进行登录").Redirect()
		ctx.Abort()
		return
	} else {
		ctx.Next()
	}
}
