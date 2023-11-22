package routers

import (
    // "fmt"
    "github.com/gin-gonic/gin"

    "FileServerWeb/views"
    "FileServerWeb/views/auth"
    "FileServerWeb/views/admin"
    "FileServerWeb/views/file"
)

func Routers(engine *gin.Engine) {
    engine.GET("/ping", views.Ping)
    engine.GET("/ip", views.IP)

    // file
    engine.POST("/upload", file.Upload)
    engine.POST("/download", file.Download)
    // engine.GET("/test", file.Test)

    // auth
    auth_group := engine.Group("/auth")
    auth_group.POST("/login", auth.Login)
    auth_group.POST("/register", auth.Register)

    // admin
    admin_group := engine.Group("/admin")
    admin_group.POST("/users_info", admin.Users_info)

}