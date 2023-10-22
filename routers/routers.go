package routers

import (
	// "fmt"
	"github.com/gin-gonic/gin"

	"FileServerWeb/views"
	"FileServerWeb/views/auth"
)

func Routers(engine *gin.Engine) {
	engine.GET("/ping", views.Ping)

	engine.POST("/upload", views.Upload)
	engine.POST("/login", auth.Login)
}