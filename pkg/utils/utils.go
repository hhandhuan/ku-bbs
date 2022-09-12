package utils

import (
	"github.com/hhandhuan/ku-bbs/pkg/utils/str"
	"github.com/hhandhuan/ku-bbs/pkg/utils/time"
	"github.com/hhandhuan/ku-bbs/pkg/utils/view"
	"html/template"
	"reflect"
)

// GetTemplateFuncMap 获取模版函数
func GetTemplateFuncMap() template.FuncMap {
	return template.FuncMap{
		"DiffForHumans":    time.DiffForHumans,
		"ToDateTimeString": time.ToDateTimeString,
		"Html":             view.Html,
		"RemindName":       view.RemindName,
		"StrLimit":         str.Limit,
	}
}

// StructToMap 结构体转换成字典
func StructToMap(s interface{}) map[string]interface{} {
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)
	data := make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}
