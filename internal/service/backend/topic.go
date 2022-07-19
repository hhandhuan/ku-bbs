package backend

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/util/gconv"
	eb "github.com/huhaophp/hblog/internal/entity/backend"
	"github.com/huhaophp/hblog/internal/model"
	"github.com/huhaophp/hblog/internal/service"
	"github.com/huhaophp/hblog/pkg/utils/page"
)

func TopicService(ctx *gin.Context) *sTopic {
	return &sTopic{ctx: service.Context(ctx)}
}

type sTopic struct {
	ctx *service.BaseContext
}

// GetList 话题列表
func (s *sTopic) GetList(req *eb.GetTopicListReq) (gin.H, error) {
	if req.Page == 0 {
		req.Page = 1
	}

	var (
		list   []*eb.Topic
		total  int64
		limit  = 20
		offset = (req.Page - 1) * limit
	)

	builder := model.Topic().M

	if len(req.Keywords) > 0 {
		builder = builder.Where("title like ?", fmt.Sprintf("%%%s%%", req.Keywords))
	}
	if len(req.UserID) > 0 {
		builder = builder.Where("user_id", req.UserID)
	}

	if c := builder.Count(&total); c.Error != nil {
		return nil, c.Error
	}

	f := builder.Preload("Publisher").Preload("Node").Limit(limit).Offset(offset).Find(&list)

	if f.Error != nil {
		return nil, f.Error
	}

	baseUrl := s.ctx.Ctx.Request.RequestURI

	pagination := page.New(int(total), limit, gconv.Int(req.Page), baseUrl)

	return gin.H{"list": list, "page": pagination, "req": req}, nil
}

// Delete 删除话题
func (s *sTopic) Delete(id int64) error {
	err := model.Topic().M.Where("id", id).Delete(&model.Topics{}).Error
	return err
}
