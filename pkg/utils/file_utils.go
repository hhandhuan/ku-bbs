package utils

func ResolvePathFileName(path string) string {
	filename := path
	for i := len(path) - 1; i > 0; i-- {
		if path[i] == '/' {
			filename = path[i+1:]
			break
		}
	}
	return filename
}
