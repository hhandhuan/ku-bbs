package frontend

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/hhandhuan/ku-bbs/internal/consts"
	fe "github.com/hhandhuan/ku-bbs/internal/entity/frontend"
	"github.com/hhandhuan/ku-bbs/internal/model"
	"github.com/hhandhuan/ku-bbs/internal/service"
	"github.com/hhandhuan/ku-bbs/pkg/utils/page"
	"gorm.io/gorm"
	"strings"
)

func TopicService(ctx *gin.Context) *sTopic {
	return &sTopic{ctx: service.Context(ctx)}
}

type sTopic struct {
	ctx *service.BaseContext
}

// Publish 发布话题
func (s *sTopic) Publish(req *fe.PublishTopicReq) (uint64, error) {
	topic := &model.Topics{
		Title:     req.Title,
		Content:   req.Content,
		NodeId:    req.NodeId,
		UserId:    s.ctx.Auth().ID,
		MDContent: req.MDContent,
	}

	if len(req.Tags) > 0 {
		topic.Tags = strings.Split(req.Tags, ",")
	}

	if r := model.Topic().M.Create(topic); r.Error != nil || r.RowsAffected <= 0 {
		return 0, errors.New("发布话题失败，请稍后再试")
	} else {
		return topic.ID, nil
	}
}

// GetDetail 获取详情
func (s *sTopic) GetDetail(topicId uint64) (*fe.Topic, error) {
	var topic *fe.Topic

	var uid uint64
	if s.ctx.Check() {
		uid = s.ctx.Auth().ID
	}

	query := model.Topic().M
	if uid > 0 {
		query = query.Preload(
			"Like",
			"user_id = ? AND source_type = ? AND state = ?",
			uid,
			consts.TopicSource,
			consts.Liked,
		)
	}

	r := query.
		Where("id", topicId).
		Preload("Publisher.Follow", func(db *gorm.DB) *gorm.DB {
			return db.Where("follows.state = ? AND follows.user_id = ?", consts.FollowedState, uid)
		}).
		Preload("Likes", func(db *gorm.DB) *gorm.DB {
			return db.Preload("User").Where("source_type = ? AND state = ?", consts.TopicSource, consts.Liked).Order("id DESC").Limit(12)
		}).
		Preload("Responder").
		Preload("Node").
		Find(&topic)
	if r.Error != nil {
		return nil, r.Error
	}

	if topic.ID <= 0 {
		return nil, errors.New("内容未找到")
	}

	data := map[string]interface{}{
		"view_count": gorm.Expr("view_count + ?", 1),
	}

	r = model.Topic().M.Where("id = ?", topicId).Updates(data)
	if r.Error != nil {
		return nil, errors.New("服务内部错误")
	}
	if r.RowsAffected <= 0 {
		return nil, errors.New("提交失败，请稍后在试")
	}

	return topic, nil
}

// GetList 获取列表
func (s *sTopic) GetList(req *fe.GetTopicListReq) (gin.H, error) {
	if req.Page == 0 {
		req.Page = 1
	}
	if req.Type == "" {
		req.Type = "reply"
	}

	var (
		topics []*fe.Topic
		total  int64
		limit  = 20
		offset = (req.Page - 1) * limit
	)

	query := model.Topic().M

	sortMap := map[string]string{"reply": "last_reply_at DESC", "latest": "created_at DESC"}
	if sort, ok := sortMap[req.Type]; ok {
		query = query.Order(sort)
	} else {
		var node *model.Nodes
		res := model.Node().M.Where("alias", req.Type).Limit(1).Find(&node)
		if res.Error != nil {
			return nil, res.Error
		}
		if node == nil {
			query = query.Where("node_id", 0)
		} else {
			query = query.Where("node_id", node.ID).Order("created_at DESC")
		}
	}

	if c := query.Count(&total); c.Error != nil {
		return nil, c.Error
	}

	f := query.Preload("Publisher").
		Preload("Node").
		Preload("Responder").
		Limit(limit).
		Offset(offset).
		Find(&topics)

	if f.Error != nil {
		return nil, f.Error
	}

	baseUrl := s.ctx.Ctx.Request.RequestURI

	pagination := page.New(int(total), limit, gconv.Int(req.Page), baseUrl)

	return gin.H{"topics": topics, "pagination": pagination, "type": req.Type}, nil
}
