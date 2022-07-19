package frontend

import (
	"context"
	"fmt"
	fe "github.com/huhaophp/hblog/internal/entity/frontend"

	"github.com/huhaophp/hblog/internal/service"
	"github.com/huhaophp/hblog/internal/service/frontend"

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

	if err := g.Validator().Data(req).Run(context.Background()); err != nil {
		s.To(to).WithError(err.FirstError()).Redirect()
		return
	}

	if id, err := frontend.CommentService(ctx).Submit(&req); err != nil {
		s.To(to).WithError(err).Redirect()
	} else {
		s.To(fmt.Sprintf("%s?j=comment%d", to, id)).WithMsg("发布成功").Redirect()
	}
}
