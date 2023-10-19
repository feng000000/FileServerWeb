package routers

import (
	// "fmt"
	"github.com/gin-gonic/gin"

	"FileServerWeb/views"
	"FileServerWeb/views/auth"
)

func Routers(r *gin.Engine) {
	r.GET("/ping", views.Ping)

	r.POST("/upload", views.Upload)
	r.POST("/login", auth.Login)
}