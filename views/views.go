package views

import (
	"fmt"
	// "log"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"FileServerWeb/config"
)

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
	  "message": "pong",
	})
}

func Upload(c *gin.Context) {
	var file, _ = c.FormFile("fuckgolang")

	fmt.Println(file.Filename)

	var dst = filepath.Join(config.FILE_PATH, file.Filename)
	// 上传文件至指定的完整文件路径
	c.SaveUploadedFile(file, dst)

	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}
