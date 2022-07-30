package frontend

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/frame/g"
	fe "github.com/hhandhuan/ku-bbs/internal/entity/frontend"
	"github.com/hhandhuan/ku-bbs/internal/service"
	"github.com/hhandhuan/ku-bbs/internal/service/frontend"
)

var Report = cReport{}

type cReport struct{}

// ReportSubmit 举报提交
func (c *cReport) ReportSubmit(ctx *gin.Context) {
	s := service.Context(ctx)

	if !s.Check() {
		s.Json(gin.H{"code": 1, "msg": "请登录后在继续操作"})
		return
	}

	var req fe.SubmitReportReq
	if err := ctx.ShouldBind(&req); err != nil {
		s.Json(gin.H{"code": 1, "msg": err.Error()})
		return
	}

	if err := g.Validator().Data(req).Run(context.Background()); err != nil {
		s.Json(gin.H{"code": 1, "msg": err.FirstError()})
		return
	}

	if err := frontend.ReportService(ctx).Store(&req); err != nil {
		s.Json(gin.H{"code": 1, "msg": err.Error()})
	} else {
		s.Json(gin.H{"code": 0, "msg": "提交举报成功"})
	}
}
