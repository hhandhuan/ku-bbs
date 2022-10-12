package frontend

import (
	"github.com/gin-gonic/gin"
	"github.com/hhandhuan/ku-bbs/internal/consts"
	"github.com/hhandhuan/ku-bbs/internal/model"
	"github.com/hhandhuan/ku-bbs/internal/service"
)

func NodeService(ctx *gin.Context) *SNode {
	return &SNode{ctx: service.Context(ctx)}
}

type SNode struct {
	ctx *service.BaseContext
}

// GetEnableNodes 获取已开启的所有节点
func (s *SNode) GetEnableNodes() ([]*model.Nodes, error) {
	var nodes []*model.Nodes
	r := model.Node().M.Where("state", consts.EnableState).Order("sort DESC").Find(&nodes)
	if r.Error != nil {
		return nil, r.Error
	} else {
		return nodes, nil
	}
}
