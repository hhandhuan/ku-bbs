package frontend

import (
	"github.com/gin-gonic/gin"
	fe "github.com/hhandhuan/ku-bbs/internal/entity/frontend"
	sv "github.com/hhandhuan/ku-bbs/internal/service"
	"github.com/hhandhuan/ku-bbs/internal/service/frontend"
)

var Notice = cNotice{}

type cNotice struct{}

func (*cNotice) HomePage(ctx *gin.Context) {
	s := sv.Context(ctx)

	if !s.Check() {
		s.To("/login").WithError("请登录后，再继续操作").Redirect()
		return
	}

	var req fe.GetRemindListReq
	if err := ctx.ShouldBind(&req); err != nil {
		s.To("/").WithError(err).Redirect()
		return
	}

	noticeService := frontend.NoticeService(ctx)

	data, err := noticeService.GetList(&req)
	if err != nil {
		s.To("/").WithError(err.Error()).Redirect()
		return
	}

	// 提醒未读数量
	remindUnread, _ := frontend.NoticeService(ctx).GetRemindUnread()
	// 系统未读数量
	systemUnread, _ := frontend.NoticeService(ctx).GetSystemUnread()

	data["remindUnread"] = remindUnread
	data["systemUnread"] = systemUnread

	// 更新未读消息状态
	noticeService.ReadAll(req.Type)

	s.View("frontend.notice.home", data)
}
