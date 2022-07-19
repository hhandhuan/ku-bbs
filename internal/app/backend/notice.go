package backend

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/frame/g"
	be "github.com/hhandhuan/ku-bbs/internal/entity/backend"
	sv "github.com/hhandhuan/ku-bbs/internal/service"
	bs "github.com/hhandhuan/ku-bbs/internal/service/backend"
	"log"
)

var Notice = cNotice{}

type cNotice struct{}

// IndexPage 消息管理
func (c *cNotice) IndexPage(ctx *gin.Context) {
	s := sv.Context(ctx)

	var req be.GetNoticeListReq
	if err := ctx.ShouldBind(&req); err != nil {
		s.Back().WithError(err).Redirect()
		return
	}

	if data, err := bs.NoticeService(ctx).GetList(&req); err != nil {
		s.Back().WithError(err).Redirect()
	} else {
		log.Println(data)
		s.View("backend.notice.index", data)
	}
}

// PublishPage 发布消息
func (c *cNotice) PublishPage(ctx *gin.Context) {
	s := sv.Context(ctx)

	s.View("backend.notice.publish", nil)
}

// PublishSubmit 发布提交
func (c *cNotice) PublishSubmit(ctx *gin.Context) {
	s := sv.Context(ctx)

	var req be.PublishNoticeReq
	if err := ctx.ShouldBind(&req); err != nil {
		s.Back().WithError(err).Redirect()
		return
	}

	if err := g.Validator().Data(req).Run(context.Background()); err != nil {
		s.Back().WithError(err.FirstError()).Redirect()
		return
	}

	if err := bs.NoticeService(ctx).Publish(&req); err != nil {
		s.Back().WithError(err).Redirect()
	} else {
		s.To("/backend/notices").WithMsg("发布成功").Redirect()
		return
	}
}
