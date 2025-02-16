package image_api

import (
	"fast_gin/global"
	"fast_gin/utils/find"
	"fast_gin/utils/md5"
	"fast_gin/utils/random"
	"fast_gin/utils/res"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"path"
	"path/filepath"
	"strings"
)

var whitelist = []string{
	".jpg",
	".jpeg",
	".png",
	".webp",
}

func (ImageApi) UploadView(c *gin.Context) {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		res.FailWithMsg("请选择文件", c)
		return
	}
	// 限制文件大小
	if fileHeader.Size > global.Config.Upload.Size*1024*1024 {
		res.FailWithMsg("文件大小超过限制", c)
		return
	}

	// 限制文件格式
	ext := strings.ToLower(filepath.Ext(fileHeader.Filename))

	if !find.InList(whitelist, ext) {
		res.FailWithMsg("文件格式不支持", c)
		return
	}

	fp := path.Join("uploads", global.Config.Upload.Dir, fileHeader.Filename)

	for {
		_, err = os.Stat(fp)
		if os.IsNotExist(err) {
			break
		}
		uploadFile, _ := fileHeader.Open()
		oldFile, _ := os.Open(fp)

		uploadFileHash := md5.MD5WithFile(uploadFile)
		oldFileHash := md5.MD5WithFile(oldFile)

		if uploadFileHash == oldFileHash {
			// 上传的图片和原图片相同
			res.Ok("/"+fp, "上传成功", c)
			return
		}
		// 上传的图片和原图片不同
		fileNameNotExt := strings.TrimSuffix(fileHeader.Filename, ext)
		newFileName := fmt.Sprintf("%s_%s%s", fileNameNotExt, random.RandStr(3), ext)
		fp = path.Join("uploads", global.Config.Upload.Dir, newFileName)
	}

	c.SaveUploadedFile(fileHeader, fp)

	res.Ok("/"+fp, "上传成功", c)
}
