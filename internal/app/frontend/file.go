package frontend

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/hhandhuan/ku-bbs/pkg/utils"

	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/hhandhuan/ku-bbs/internal/service"
	"github.com/hhandhuan/ku-bbs/pkg/config"
)

var File = cFile{}

type cFile struct{}

// MDUploadSubmit markdown 文件上传
func (c *cFile) MDUploadSubmit(ctx *gin.Context) {
	s := service.Context(ctx)

	file, err := ctx.FormFile("editormd-image-file")
	if err != nil {
		s.MDFileJson(0, err.Error(), "")
		return
	}

	// 目前限制 M 图片大小
	if file.Size > 1024*1024*config.GetInstance().Upload.Filesize {
		s.MDFileJson(0, "仅支持小于 2M 大小的图片", "")
		return
	}

	arr := strings.Split(file.Filename, ".")
	ext := arr[len(arr)-1]

	// 检查图片格式
	if !gstr.InArray(config.GetInstance().Upload.Ext, ext) {
		s.MDFileJson(0, "file format not supported", "")
		return
	}

	path := fmt.Sprintf("%s/topic", config.GetInstance().Upload.Path)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, os.ModePerm)
		os.Chmod(path, os.ModePerm)
	}

	name := utils.Md5(time.Now().String()+file.Filename) + "." + ext

	err = ctx.SaveUploadedFile(file, fmt.Sprintf("%s/%s", path, name))
	if err != nil {
		s.MDFileJson(0, err.Error(), "")
	} else {
		s.MDFileJson(1, "ok", fmt.Sprintf("/assets/upload/topic/%s", name))
	}
}
