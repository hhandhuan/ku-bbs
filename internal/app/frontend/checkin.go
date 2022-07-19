package frontend

import (
	"github.com/gin-gonic/gin"
	"github.com/hhandhuan/ku-bbs/internal/service"
	"github.com/hhandhuan/ku-bbs/internal/service/frontend"
)

var Checkin = cCheckin{}

type cCheckin struct{}

// StoreSubmit 提交评论
func (c *cCheckin) StoreSubmit(ctx *gin.Context) {
	s := service.Context(ctx)

	if !s.Check() {
		s.Json(gin.H{"code": 1, "msg": "请登录后在继续操作"})
		return
	}

	if err := frontend.CheckinService(ctx).Store(); err != nil {
		s.Json(gin.H{"code": 1, "msg": err.Error()})
	} else {
		s.Json(gin.H{"code": 0, "msg": "ok"})
	}
}
