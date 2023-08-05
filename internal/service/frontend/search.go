package frontend

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/util/gconv"
	fe "github.com/hhandhuan/ku-bbs/internal/entity/frontend"
	"github.com/hhandhuan/ku-bbs/internal/model"
	"github.com/hhandhuan/ku-bbs/internal/service"
	"github.com/hhandhuan/ku-bbs/pkg/utils/page"
)

type SSearch struct{ ctx *service.BaseContext }

func SearchService(ctx *gin.Context) *SSearch {
	return &SSearch{ctx: service.Context(ctx)}
}

// GetList 获取搜索列表
func (s *SSearch) GetList(req *fe.GetSearchListReq) (gin.H, error) {
	if req.Page == 0 {
		req.Page = 1
	}

	var (
		topics []*fe.Topic
		total  int64
		limit  = 20
		offset = (req.Page - 1) * limit
	)

	query := model.Topic()

	if len(req.Keywords) > 0 {
		key := fmt.Sprintf("%%%s%%", req.Keywords)
		query = query.Where("title like ?", key).Or("content like ?", key)
	} else {
		return gin.H{"list": topics, "page": nil, "req": req, "total": total}, nil
	}

	if c := query.Count(&total); c.Error != nil {
		return nil, c.Error
	}

	f := query.
		Preload("Publisher").
		Preload("Node").
		Preload("Responder").
		Order("id DESC").
		Limit(limit).
		Offset(offset).
		Find(&topics)

	if f.Error != nil {
		return nil, f.Error
	}

	pageObj := page.New(int(total), limit, gconv.Int(req.Page), s.ctx.Ctx.Request.RequestURI)

	return gin.H{"list": topics, "page": pageObj, "req": req, "total": total}, nil
}
