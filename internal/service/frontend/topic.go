package frontend

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"unicode/utf8"

	"github.com/hhandhuan/ku-bbs/pkg/utils/str"

	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/hhandhuan/ku-bbs/internal/consts"
	fe "github.com/hhandhuan/ku-bbs/internal/entity/frontend"
	"github.com/hhandhuan/ku-bbs/internal/model"
	"github.com/hhandhuan/ku-bbs/internal/service"
	"github.com/hhandhuan/ku-bbs/pkg/utils/page"
	time2 "github.com/hhandhuan/ku-bbs/pkg/utils/time"
	"gorm.io/gorm"
)

const MaxTagsLen = 3 // 标签最大长度
const MaxTagLen = 15 // 单个标签最大长度

func TopicService(ctx *gin.Context) *STopic {
	return &STopic{ctx: service.Context(ctx)}
}

type STopic struct {
	ctx *service.BaseContext
}

// Publish 发布话题
func (s *STopic) Publish(req *fe.PublishTopicReq) (uint64, error) {
	brief, err := s.ParseBrief(req.Content)
	if err != nil {
		return 0, err
	}
	images, err := s.ParseHtmlImages(req.Content)
	if err != nil {
		return 0, err
	}

	data := &model.Topics{
		Title:        req.Title,
		Content:      req.Content,
		NodeId:       req.NodeId,
		UserId:       s.ctx.Auth().ID,
		MDContent:    req.MDContent,
		CommentState: consts.EnableState,
		Brief:        str.Limit(brief, 0, 100, "..."),
		Images:       images,
	}

	// 检查话题标签
	tags := strings.Split(req.Tags, ",")
	if len(tags) > 0 {
		if len(tags) > MaxTagsLen {
			return 0, fmt.Errorf("最多添加%d标签", MaxTagsLen)
		}
		isOk := true
		for _, value := range tags {
			if utf8.RuneCountInString(value) > MaxTagLen {
				isOk = false
				break
			}
		}
		if !isOk {
			return 0, fmt.Errorf("单个标签最多%d个字符", MaxTagLen)
		} else {
			data.Tags = tags
		}
	}

	r := model.Topic().Create(data)
	if r.Error != nil || r.RowsAffected <= 0 {
		return 0, errors.New("发布话题失败，请稍后再试")
	} else {
		return data.ID, nil
	}
}

// GetDetail 获取详情
func (s *STopic) GetDetail(topicId uint64) (*fe.Topic, error) {
	var topic *fe.Topic

	var uid uint64
	if s.ctx.Check() {
		uid = s.ctx.Auth().ID
	}

	query := model.Topic()
	if uid > 0 {
		query = query.Preload("Like", "user_id = ? AND source_type = ? AND state = ?", uid, consts.TopicSource, consts.Liked)
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
		return nil, errors.New("话题已删除或不存在")
	}

	data := map[string]interface{}{
		"view_count": gorm.Expr("view_count + ?", 1),
	}

	r = model.Topic().Where("id = ?", topicId).Updates(data)
	if r.Error != nil {
		return nil, errors.New("服务内部错误")
	}
	if r.RowsAffected <= 0 {
		return nil, errors.New("提交失败，请稍后在试")
	}

	topic.PostDays = time2.DiffDays(topic.CreatedAt)

	return topic, nil
}

// GetList 获取列表
func (s *STopic) GetList(req *fe.GetTopicListReq) (gin.H, error) {
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

	query := model.Topic()

	sortMap := map[string]string{
		"reply":  "type DESC,last_reply_at DESC",
		"latest": "type DESC,created_at DESC",
		"node":   "type DESC,created_at DESC",
	}

	if sort, ok := sortMap[req.Type]; ok {
		query = query.Order(sort)
	} else {
		var node *model.Nodes
		res := model.Node().Where("alias", req.Type).Limit(1).Find(&node)
		if res.Error != nil {
			return nil, res.Error
		}
		if node == nil {
			query = query.Where("node_id", 0)
		} else {
			query = query.Where("node_id", node.ID).Order(sortMap["node"])
		}
	}

	if c := query.Count(&total); c.Error != nil {
		return nil, c.Error
	}

	f := query.Preload("Publisher").
		Preload("Node", "state = ?", consts.EnableState).
		Preload("Responder").
		Limit(limit).
		Offset(offset).
		Find(&topics)

	if f.Error != nil {
		return nil, f.Error
	}

	baseUrl := s.ctx.Ctx.Request.RequestURI

	pageObj := page.New(int(total), limit, gconv.Int(req.Page), baseUrl)

	return gin.H{"topics": topics, "page": pageObj, "type": req.Type}, nil
}

// Delete 删除话题
func (s *STopic) Delete(ID uint64) error {
	if !s.ctx.Check() {
		return errors.New("请登录后在继续操作")
	}

	var (
		topic   *model.Topics
		comment *model.Comments
	)

	// 检查话题下是否存在评论
	f := model.Comment().Unscoped().Where("topic_id = ?", ID).Find(&comment)
	if f.Error != nil {
		log.Println("delete topic error: ", f.Error)
		return f.Error
	}
	if comment.ID > 0 {
		return errors.New("话题下存在评论，无法删除")
	}

	// 检查话题是否存在
	f = model.Topic().Where("id = ?", ID).Find(&topic)
	if f.Error != nil {
		log.Println("delete topic error: ", f.Error)
		return f.Error
	}
	if topic.ID <= 0 {
		return errors.New("资源未找到")
	}

	// 检查权限
	if s.ctx.Auth().ID != topic.UserId {
		return errors.New("无权限操作")
	}

	// 删除话题
	d := model.Topic().Delete(&model.Topics{}, ID)
	if d.Error != nil {
		log.Println("delete topic error: ", d.Error)
		return f.Error
	}
	if d.RowsAffected <= 0 {
		return errors.New("目标已删除或不存在")
	}

	return nil
}

func (s *STopic) Edit(ID uint64, req *fe.PublishTopicReq) (uint64, error) {
	if !s.ctx.Check() {
		return 0, errors.New("请登录后在继续操作")
	}

	var topic *model.Topics
	// 检查话题是否存在
	f := model.Topic().Where("id = ?", ID).Find(&topic)
	if f.Error != nil {
		log.Println("delete topic error: ", f.Error)
		return 0, f.Error
	}
	if topic.ID <= 0 {
		return 0, errors.New("资源未找到")
	}

	// 检查权限
	if s.ctx.Auth().ID != topic.UserId {
		return 0, errors.New("无权限操作")
	}

	brief, err := s.ParseBrief(req.Content)
	if err != nil {
		return 0, err
	}
	images, err := s.ParseHtmlImages(req.Content)
	if err != nil {
		return 0, err
	}

	data := &model.Topics{
		Title:     req.Title,
		Content:   req.Content,
		NodeId:    req.NodeId,
		MDContent: req.MDContent,
		Brief:     str.Limit(brief, 0, 100, "..."),
		Images:    images,
	}

	log.Println(req.Tags)

	// 检查话题标签
	tags := strings.Split(req.Tags, ",")
	if len(tags) > 0 {
		if len(tags) > MaxTagsLen {
			return 0, errors.New(fmt.Sprintf("最多添加%d标签", MaxTagsLen))
		}
		isOk := true
		for _, value := range tags {
			if utf8.RuneCountInString(value) > MaxTagLen {
				isOk = false
				break
			}
		}
		if !isOk {
			return 0, errors.New(fmt.Sprintf("单个标签最多%d个字符", MaxTagLen))
		} else {
			data.Tags = tags
		}
	}

	r := model.Topic().Where("id = ?", ID).Updates(data)
	if r.Error != nil || r.RowsAffected <= 0 {
		return 0, errors.New("编辑话题失败，请稍后再试")
	} else {
		return topic.ID, nil
	}
}

// SettingCommentState 设置话题讨论状态
func (s *STopic) SettingCommentState(ID uint64) error {
	if !s.ctx.Check() {
		return errors.New("请登录后在继续操作")
	}

	var topic *model.Topics
	// 检查话题是否存在
	f := model.Topic().Where("id = ?", ID).Find(&topic)
	if f.Error != nil {
		log.Println("delete topic error: ", f.Error)
		return f.Error
	}
	if topic.ID <= 0 {
		return errors.New("话题资源未找到")
	}
	// 检查权限
	if s.ctx.Auth().ID != topic.UserId {
		return errors.New("无权限操作")
	}

	state := consts.EnableState
	if topic.CommentState == consts.EnableState {
		state = consts.DisableState
	}
	// 更新评论状态
	d := model.Topic().Where("id", ID).Update("comment_state", state)
	if d.Error != nil {
		log.Println("update topic state error: ", d.Error)
		return f.Error
	}
	if d.RowsAffected <= 0 {
		return errors.New("目标已删除或不存在")
	}

	return nil
}

// ParseBrief parse topic brief
func (s *STopic) ParseBrief(html string) (string, error) {
	if doc, err := goquery.NewDocumentFromReader(strings.NewReader(html)); err != nil {
		return "", fmt.Errorf("html to text error: %v", err)
	} else {
		doc.Find("img").ReplaceWithHtml("[图片]")
		rs := strings.Replace(doc.Text(), " ", "", -1)
		rs = strings.Replace(rs, "\n", "", -1)
		return rs, nil
	}
}

// ParseHtmlImages parse topic images
func (s *STopic) ParseHtmlImages(html string) ([]string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return nil, fmt.Errorf("html to text error: %v", err)
	}
	images := make([]string, 0)
	doc.Find("img").Each(func(i int, selection *goquery.Selection) {
		img, _ := selection.Attr("src")
		images = append(images, img)
	})
	return images, nil
}
