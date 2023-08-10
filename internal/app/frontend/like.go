package frontend

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/frame/g"
	fe "github.com/hhandhuan/ku-bbs/internal/entity/frontend"
	"github.com/hhandhuan/ku-bbs/internal/service"
	"github.com/hhandhuan/ku-bbs/internal/service/frontend"
)

var Like = cLike{}

type cLike struct{}

// LikeSubmit 点赞提交
func (c *cLike) LikeSubmit(ctx *gin.Context) {
	s := service.Context(ctx)

	if !s.Check() {
		s.Json(gin.H{"code": 1, "msg": "请登录后在继续操作"})
		return
	}

	var req fe.LikeReq
	if err := ctx.ShouldBind(&req); err != nil {
		s.Json(gin.H{"code": 1, "msg": err.Error()})
		return
	}

	err := g.Validator().Data(req).Run(context.Background())
	if err != nil {
		s.Json(gin.H{"code": 1, "msg": err.FirstError()})
		return
	}

	lerr := frontend.LikeService(ctx).Like(&req)
	if lerr != nil {
		s.Json(gin.H{"code": 1, "msg": lerr.Error()})
	} else {
		s.Json(gin.H{"code": 0, "msg": "ok"})
	}
}
