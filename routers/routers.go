package routers

import (
	// "fmt"
	"github.com/gin-gonic/gin"

	"FileServerWeb/views"
)

func Routers(r *gin.Engine) {
	r.GET("/ping", views.Ping)

	r.POST("/upload", views.Upload)
}