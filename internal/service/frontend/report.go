package frontend

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/hhandhuan/ku-bbs/internal/consts"
	fe "github.com/hhandhuan/ku-bbs/internal/entity/frontend"
	"github.com/hhandhuan/ku-bbs/internal/model"
	"github.com/hhandhuan/ku-bbs/internal/service"
)

func ReportService(ctx *gin.Context) *SReport {
	return &SReport{ctx: service.Context(ctx)}
}

type SReport struct {
	ctx *service.BaseContext
}

func (s *SReport) Store(req *fe.SubmitReportReq) error {
	if req.SourceType == consts.ReportTopicSource {
		var topic *model.Topics
		f := model.Topic().M.Where("id", req.TargetID).Find(&topic)
		if f.Error != nil {
			return fmt.Errorf("服务内部错误: %v", f.Error)
		}
		if topic == nil {
			return errors.New("举报目标不存在")
		}

		var report *model.Reports
		f = model.Report().M.Where(map[string]interface{}{
			"source_id":   req.SourceID,
			"source_type": req.SourceType,
			"target_id":   req.TargetID,
			"user_id":     s.ctx.Auth().ID,
		}).Find(&report)
		if f.Error != nil {
			return fmt.Errorf("服务内部错误: %v", f.Error)
		}
		if report.ID > 0 {
			return errors.New("举报已存在")
		}

		c := model.Report().M.Create(&model.Reports{
			UserId:     s.ctx.Auth().ID,
			Remark:     req.Remark,
			SourceId:   req.SourceID,
			SourceType: req.SourceType,
			TargetId:   req.TargetID,
			SourceUrl:  fmt.Sprintf("/topics/%d", req.SourceID),
			State:      consts.ReportAwaitingState,
		})
		if c.Error != nil && c.RowsAffected <= 0 {
			return fmt.Errorf("服务内部错误: %v", f.Error)
		}
		return nil
	} else {
		var comment *model.Comments
		f := model.Comment().M.Where("id", req.SourceID).Find(&comment)
		if f.Error != nil {
			return fmt.Errorf("服务内部错误: %v", f.Error)
		}
		if comment == nil {
			return errors.New("举报目标不存在")
		}

		var report *model.Reports
		f = model.Report().M.Where(map[string]interface{}{
			"source_id":   req.SourceID,
			"source_type": req.SourceType,
			"target_id":   req.TargetID,
			"user_id":     s.ctx.Auth().ID,
		}).Find(&report)
		if f.Error != nil {
			return fmt.Errorf("服务内部错误: %v", f.Error)
		}
		if report.ID > 0 {
			return errors.New("举报已存在")
		}

		c := model.Report().M.Create(&model.Reports{
			UserId:     s.ctx.Auth().ID,
			Remark:     req.Remark,
			SourceId:   req.SourceID,
			SourceType: req.SourceType,
			TargetId:   req.TargetID,
			SourceUrl:  fmt.Sprintf("/topics/%d?j=comment%d", comment.TopicId, comment.ID),
			State:      consts.ReportAwaitingState,
		})
		if c.Error != nil && c.RowsAffected <= 0 {
			return fmt.Errorf("服务内部错误: %v", f.Error)
		}
	}

	return nil
}
