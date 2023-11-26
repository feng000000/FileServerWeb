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
    file_group := engine.Group("/file")
    {
        file_group.POST("/upload", file.Upload)
        file_group.POST("/download", file.Download)
        file_group.POST("/storage_usage", file.StorageUsage)

    }

    // auth
    auth_group := engine.Group("/auth")
    {
        auth_group.POST("/login", auth.Login)
        auth_group.POST("/register", auth.Register)
    }

    // admin
    admin_group := engine.Group("/admin")
    {
        admin_group.POST("/users_info", admin.Users_info)
    }

}