package file

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"FileServerWeb/config"
)

func Upload(c *gin.Context) {
	var file, _ = c.FormFile("fuckgolang")

	fmt.Println(file.Filename)

	var dst = filepath.Join(config.DATA_FILE_PATH, file.Filename)
	// 上传文件至指定的完整文件路径
	c.SaveUploadedFile(file, dst)

	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}
