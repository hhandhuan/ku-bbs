package view

import "html/template"

// Html 解析 HTML
func Html(s string) template.HTML {
	return template.HTML(s)
}

// RemindName 提醒名
func RemindName(a string) string {
	m := map[string]string{
		"comment:topic": "评论了你的话题",
		"reply:comment": "回复了你的评论",
		"like:topic":    "赞了你的话题",
		"like:comment":  "赞了你的评论",
		"follow:user":   "关注了你",
	}
	if v, ok := m[a]; !ok {
		return ""
	} else {
		return v
	}
}
