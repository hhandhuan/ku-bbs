package utils

import (
	"github.com/huhaophp/hblog/pkg/utils/time"
	"github.com/huhaophp/hblog/pkg/utils/view"
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
