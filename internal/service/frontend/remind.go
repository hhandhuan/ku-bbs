package frontend

import (
	"github.com/gin-gonic/gin"
	"github.com/hhandhuan/ku-bbs/internal/service"
)

func RemindService(ctx *gin.Context) *SRemind {
	return &SRemind{ctx: service.Context(ctx)}
}

type SRemind struct {
	ctx *service.BaseContext
}
