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

	// auth
	auth_group := engine.Group("/auth")
	auth_group.POST("/login", auth.Login)
	auth_group.POST("/register", auth.Register)
}