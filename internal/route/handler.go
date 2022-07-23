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
