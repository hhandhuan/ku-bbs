package frontend

import (
	"context"
	"fmt"

	fe "github.com/hhandhuan/ku-bbs/internal/entity/frontend"
	"github.com/hhandhuan/ku-bbs/internal/service"
	"github.com/hhandhuan/ku-bbs/internal/service/frontend"

	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/frame/g"
)

var Comment = cComment{}

type cComment struct{}

// PublishSubmit 提交评论
func (c *cComment) PublishSubmit(ctx *gin.Context) {
	s := service.Context(ctx)

	var req fe.SubmitCommentReq
	if err := ctx.ShouldBind(&req); err != nil {
		s.Back().WithError(err).Redirect()
		return
	}

	to := fmt.Sprintf("/topics/%d", req.TopicId)

	verr := g.Validator().Data(req).Run(context.Background())
	if verr != nil {
		s.To(to).WithError(verr.FirstError()).Redirect()
		return
	}

	id, err := frontend.CommentService(ctx).Submit(&req)
	if err != nil {
		s.To(to).WithError(err).Redirect()
	} else {
		s.To(fmt.Sprintf("%s?j=comment%d", to, id)).WithMsg("发布成功").Redirect()
	}
}

// DeleteSubmit 删除评论
func (c *cComment) DeleteSubmit(ctx *gin.Context) {
	s := service.Context(ctx)

	var req fe.DeleteCommentReq
	if err := ctx.ShouldBind(&req); err != nil {
		s.Back().WithError(err).Redirect()
		return
	}

	verr := g.Validator().Data(req).Run(context.Background())
	if verr != nil {
		s.Json(gin.H{"code": 1, "msg": verr.FirstError()})
		return
	}

	derr := frontend.CommentService(ctx).Delete(req.ID)
	if derr != nil {
		s.Json(gin.H{"code": 1, "msg": "删除失败"})
	} else {
		s.Json(gin.H{"code": 0, "msg": "删除成功"})
	}
}
