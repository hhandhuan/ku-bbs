package backend

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/util/gconv"
	eb "github.com/hhandhuan/ku-bbs/internal/entity/backend"
	"github.com/hhandhuan/ku-bbs/internal/model"
	"github.com/hhandhuan/ku-bbs/internal/service"
	"github.com/hhandhuan/ku-bbs/pkg/utils/page"
	"log"
)

func NoticeService(ctx *gin.Context) *sNotice {
	return &sNotice{ctx: service.Context(ctx)}
}

type sNotice struct {
	ctx *service.BaseContext
}

// GetList 消息列表
func (s *sNotice) GetList(req *eb.GetNoticeListReq) (gin.H, error) {
	if req.Page == 0 {
		req.Page = 1
	}

	var (
		list   []*eb.SystemNotices
		total  int64
		limit  = 20
		offset = (req.Page - 1) * limit
	)

	builder := model.SystemNotice().M

	if len(req.Keywords) > 0 {
		builder = builder.Where("title like ?", fmt.Sprintf("%%%s%%", req.Keywords))
	}

	if c := builder.Count(&total); c.Error != nil {
		return nil, c.Error
	}

	f := builder.Preload("Publisher").Preload("TargetUser").Limit(limit).Offset(offset).Find(&list)

	if f.Error != nil {
		return nil, f.Error
	}

	pageObj := page.New(int(total), limit, gconv.Int(req.Page), s.ctx.Ctx.Request.RequestURI)

	return gin.H{"list": list, "page": pageObj, "req": req}, nil
}

// Publish 发布消息
func (s *sNotice) Publish(req *eb.PublishNoticeReq) error {
	notice := &model.SystemNotices{
		UserId:    s.ctx.Auth().ID,
		Title:     req.Title,
		TargetId:  req.TargetId,
		Content:   req.Content,
		MDContent: req.MDContent,
	}
	if err := model.SystemNotice().M.Create(notice).Error; err != nil {
		return err
	}
	if notice.ID <= 0 {
		return errors.New("消息发布失败")
	}

	var users []*model.Users
	f := model.User().M.Select("id").Find(&users)
	if f.Error != nil || users == nil {
		log.Println(f.Error)
		return errors.New("消息发布失败")
	}
	if users == nil {

	}

	var userNotices []model.SystemUserNotices
	for _, value := range users {
		userNotices = append(userNotices, model.SystemUserNotices{
			UserId:   value.ID,
			NoticeId: notice.ID,
		})
	}

	if err := model.SystemUserNotice().M.Create(&userNotices).Error; err != nil {
		return f.Error
	}

	return nil
}
