package utils

import (
	"github.com/hhandhuan/ku-bbs/pkg/utils/time"
	"github.com/hhandhuan/ku-bbs/pkg/utils/view"
	"html/template"
)

// GetTemplateFuncMap 获取模版函数
func GetTemplateFuncMap() template.FuncMap {
	return template.FuncMap{
		"DiffForHumans": time.DiffForHumans,
		"FormatTime":    time.FormatTime,
		"Html":          view.Html,
		"RemindName":    view.RemindName,
	}
}
