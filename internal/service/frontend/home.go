package frontend

import (
	"github.com/gin-gonic/gin"
	"github.com/hhandhuan/ku-bbs/internal/service"
)

func HomeService(ctx *gin.Context) *sHome {
	return &sHome{ctx: service.Context(ctx)}
}

type sHome struct {
	ctx *service.BaseContext
}
