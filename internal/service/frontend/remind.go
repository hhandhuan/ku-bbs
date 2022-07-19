package frontend

import (
	"github.com/gin-gonic/gin"
	"github.com/hhandhuan/ku-bbs/internal/service"
)

func RemindService(ctx *gin.Context) *sRemind {
	return &sRemind{ctx: service.Context(ctx)}
}

type sRemind struct {
	ctx *service.BaseContext
}
