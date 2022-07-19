package frontend

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	fe "github.com/hhandhuan/ku-bbs/internal/entity/frontend"
	"github.com/hhandhuan/ku-bbs/internal/service"
	"github.com/hhandhuan/ku-bbs/internal/service/frontend"
)

var Topic = cTopic{}

type cTopic struct{}

// PublishPage 发布页面
func (c *cTopic) PublishPage(ctx *gin.Context) {
	s := service.Context(ctx)

	if !s.Check() {
		s.To("/login").WithError("请登录后，再继续操作").Redirect()
		return
	}

	if nodes, err := frontend.NodeService(ctx).GetEnableNodes(); err != nil {
		s.To("/").WithError(err.Error()).Redirect()
	} else {
		s.View("frontend.topic.publish", gin.H{"nodes": nodes})
	}
}

// PublishSubmit 发布提交
func (c *cTopic) PublishSubmit(ctx *gin.Context) {
	s := service.Context(ctx)

	var req fe.PublishTopicReq
	if err := ctx.ShouldBind(&req); err != nil {
		s.Back().WithError(err).Redirect()
		return
	}

	if err := g.Validator().Data(req).Run(context.Background()); err != nil {
		s.Back().WithError(err.FirstError()).Redirect()
		return
	}

	if id, err := frontend.TopicService(ctx).Publish(&req); err != nil {
		s.Back().WithError(err).Redirect()
	} else {
		s.To(fmt.Sprintf("/topics/%d", id)).WithMsg("发布成功").Redirect()
	}
}

// DetailPage 话题详情
func (c *cTopic) DetailPage(ctx *gin.Context) {
	s := service.Context(ctx)

	id := gconv.Uint64(ctx.Param("id"))

	topic, err := frontend.TopicService(ctx).GetDetail(id)
	if err != nil {
		s.To("/").WithError(err).Redirect()
		return
	}

	items, err := frontend.CommentService(ctx).GetList(id)
	if err != nil {
		s.To("/").WithError(err).Redirect()
		return
	}

	s.View("frontend.topic.detail", gin.H{"topic": topic, "comments": items})
}
