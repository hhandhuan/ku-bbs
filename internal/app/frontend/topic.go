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

	nodes, err := frontend.NodeService(ctx).GetEnableNodes()
	if err != nil {
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
		s.Back().WithError(err).WithData(req).Redirect()
		return
	}

	verr := g.Validator().Data(req).Run(context.Background())
	if verr != nil {
		s.Back().WithError(verr.FirstError()).WithData(req).Redirect()
		return
	}

	id, err := frontend.TopicService(ctx).Publish(&req)
	if err != nil {
		s.Back().WithError(err).WithData(req).Redirect()
	} else {
		s.To(fmt.Sprintf("/topics/%d", id)).WithMsg("发布成功").Redirect()
	}
}

// DetailPage 话题详情
func (c *cTopic) DetailPage(ctx *gin.Context) {
	s := service.Context(ctx)

	topicID := gconv.Uint64(ctx.Param("id"))
	authorID := gconv.Uint64(ctx.Query("author_id"))

	topic, err := frontend.TopicService(ctx).GetDetail(topicID)
	if err != nil {
		s.To("/").WithError(err).Redirect()
		return
	}

	list, err := frontend.CommentService(ctx).GetList(topicID, authorID)
	if err != nil {
		s.To("/").WithError(err).Redirect()
		return
	}

	data := gin.H{"topic": topic, "comments": list, "author_id": authorID}

	s.SetTitle(topic.Title).View("frontend.topic.detail", data)
}

// DeleteSubmit 删除话题
func (c *cTopic) DeleteSubmit(ctx *gin.Context) {
	s := service.Context(ctx)
	i := gconv.Uint64(ctx.Param("id"))

	err := frontend.TopicService(ctx).Delete(i)
	if err != nil {
		s.To(fmt.Sprintf("/topics/%d", i)).WithError(err).Redirect()
	} else {
		s.To("/").WithMsg("删除成功").Redirect()
	}
}

// EditPage 编辑话题
func (c *cTopic) EditPage(ctx *gin.Context) {
	s := service.Context(ctx)
	i := gconv.Uint64(ctx.Param("id"))

	if !s.Check() {
		s.To("/login").WithError("请登录后，再继续操作").Redirect()
		return
	}

	nodes, err := frontend.NodeService(ctx).GetEnableNodes()
	if err != nil {
		s.To("/").WithError(err.Error()).Redirect()
		return
	}

	topic, err := frontend.TopicService(ctx).GetDetail(i)
	if err != nil {
		s.To("/").WithError(err.Error()).Redirect()
		return
	}

	s.View("frontend.topic.edit", gin.H{"nodes": nodes, "topic": topic})
}

// EditSubmit 编辑提交
func (c *cTopic) EditSubmit(ctx *gin.Context) {
	s := service.Context(ctx)
	i := gconv.Uint64(ctx.Param("id"))

	var req fe.PublishTopicReq
	if err := ctx.ShouldBind(&req); err != nil {
		s.Back().WithError(err).WithData(req).Redirect()
		return
	}

	verr := g.Validator().Data(req).Run(context.Background())
	if verr != nil {
		s.Back().WithError(verr.FirstError()).WithData(req).Redirect()
		return
	}

	id, err := frontend.TopicService(ctx).Edit(i, &req)
	if err != nil {
		s.Back().WithError(err).WithData(req).Redirect()
	} else {
		s.To(fmt.Sprintf("/topics/%d", id)).WithMsg("编辑成功").Redirect()
	}
}

// SettingCommentStateSubmit 设置话题讨论状态
func (c *cTopic) SettingCommentStateSubmit(ctx *gin.Context) {
	s := service.Context(ctx)
	i := gconv.Uint64(ctx.Param("id"))

	err := frontend.TopicService(ctx).SettingCommentState(i)
	if err != nil {
		s.To(fmt.Sprintf("/topics/%d", i)).WithError(err).Redirect()
	} else {
		s.To(fmt.Sprintf("/topics/%d", i)).WithMsg("操作成功").Redirect()
	}
}
