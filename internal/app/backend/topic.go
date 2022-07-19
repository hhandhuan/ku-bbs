package backend

import (
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/util/gconv"
	be "github.com/hhandhuan/ku-bbs/internal/entity/backend"
	sv "github.com/hhandhuan/ku-bbs/internal/service"
	bs "github.com/hhandhuan/ku-bbs/internal/service/backend"
	"log"
)

var Topic = cTopic{}

type cTopic struct{}

// IndexPage 用户主页
func (c *cTopic) IndexPage(ctx *gin.Context) {
	s := sv.Context(ctx)

	var req be.GetTopicListReq
	if err := ctx.ShouldBind(&req); err != nil {
		s.Back().WithError(err).Redirect()
		return
	}

	data, err := bs.TopicService(ctx).GetList(&req)
	if err != nil {
		s.Back().WithError(err).Redirect()
	} else {
		s.View("backend.topic.index", data)
	}
}

// DeleteSubmit 删除提交
func (c *cTopic) DeleteSubmit(ctx *gin.Context) {
	s := sv.Context(ctx)
	t := "/backend/topics"

	id := gconv.Int64(ctx.Param("id"))
	if id <= 0 {
		s.To(t).WithError("参数错误").Redirect()
		return
	}

	if err := bs.TopicService(ctx).Delete(id); err != nil {
		log.Println(err)
		s.To(t).WithError("删除失败").Redirect()
	} else {
		s.To(t).WithMsg("删除成功").Redirect()
	}
}
