package frontend

import (
	"github.com/gin-gonic/gin"
	"github.com/hhandhuan/ku-bbs/internal/service"
)

func HomeService(ctx *gin.Context) *SHome {
	return &SHome{ctx: service.Context(ctx)}
}

type SHome struct {
	ctx *service.BaseContext
}
