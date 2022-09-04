package str

import "unicode/utf8"

// Substr 截取字符串
func Substr(s string, start, length int) string {
	strLen := utf8.RuneCountInString(s)
	if strLen <= 0 {
		return ""
	}
	if length > strLen {
		length = strLen
	}

	runes := []rune(s)

	return string(runes[start:length])
}

// Limit 将字符串以指定长度进行截断
func Limit(s string, start, length int, append string) string {
	strLen := utf8.RuneCountInString(s)
	if strLen <= 0 {
		return ""
	}

	if length >= strLen {
		length = strLen
		append = ""
	}

	runes := []rune(s)

	return string(runes[start:length]) + append
}
