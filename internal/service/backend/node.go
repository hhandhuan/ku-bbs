package backend

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	eb "github.com/hhandhuan/ku-bbs/internal/entity/backend"
	"github.com/hhandhuan/ku-bbs/internal/model"
	"github.com/hhandhuan/ku-bbs/internal/service"
)

func NodeService(ctx *gin.Context) *sNode {
	return &sNode{ctx: service.Context(ctx)}
}

type sNode struct {
	ctx *service.BaseContext
}

// GetList 消息列表
func (s *sNode) GetList(req *eb.GetNodeListReq) (gin.H, error) {
	var list []*model.Nodes

	builder := model.Node()
	if len(req.Keywords) > 0 {
		builder = builder.Where("name like ?", fmt.Sprintf("%%%s%%", req.Keywords))
	}

	if f := builder.Order("sort DESC").Find(&list); f.Error != nil {
		return nil, f.Error
	}

	return gin.H{"list": list, "req": req}, nil
}

// Create 创建节点
func (s *sNode) Create(req *eb.CreateNodeReq) error {
	var node *model.Nodes
	f := model.Node().Where("name", req.Name).Or("alias", req.Alias).Find(&node)
	if f.Error != nil {
		return f.Error
	}
	if node.ID > 0 {
		return errors.New("节点已存在，无法重复创建")
	}
	if c := model.Node().Create(&model.Nodes{
		Name:  req.Name,
		Alias: req.Alias,
		Sort:  req.Sort,
		State: req.State,
		Desc:  req.Desc,
	}); c.Error != nil || c.RowsAffected <= 0 {
		return c.Error
	} else {
		return nil
	}
}

// Edit 编辑节点
func (s *sNode) Edit(id uint64, req *eb.CreateNodeReq) error {
	var node *model.Nodes
	f := model.Node().Where("id != ? AND (name = ? OR alias = ?)", id, req.Name, req.Alias).Find(&node)
	if f.Error != nil {
		return f.Error
	}
	if node.ID > 0 {
		return errors.New("节点已存在，无法重复创建")
	}

	if c := model.Node().Where("id = ?", id).Updates(map[string]interface{}{
		"name":  req.Name,
		"alias": req.Alias,
		"sort":  req.Sort,
		"state": req.State,
		"desc":  req.Desc,
	}); c.Error != nil || c.RowsAffected <= 0 {
		return c.Error
	} else {
		return nil
	}
}

// GetDetail 获取详情
func (s *sNode) GetDetail(id uint64) (*model.Nodes, error) {
	var node *model.Nodes
	c := model.Node().Where("id", id).Find(&node)
	if c.Error != nil {
		return nil, c.Error
	}
	if node.ID <= 0 {
		return nil, errors.New("节点未找到")
	} else {
		return node, nil
	}
}

// Delete 删除节点
func (s *sNode) Delete(id int64) error {
	err := model.Node().Where("id", id).Delete(&model.Nodes{}).Error
	return err
}
