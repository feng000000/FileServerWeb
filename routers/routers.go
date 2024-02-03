package routers

import (
    // "fmt"
    "github.com/gin-gonic/gin"

    "FileServerWeb/middleware"
    "FileServerWeb/views"
    "FileServerWeb/views/admin"
    "FileServerWeb/views/auth"
    "FileServerWeb/views/file"
    "FileServerWeb/views/user"
)

func Routers(engine *gin.Engine) {
    engine.GET("/ping", views.Ping)
    engine.GET("/ip", views.IP)

    // auth
    var authGroup = engine.Group("/auth")
    {
        authGroup.POST("/login", auth.LoginHandler)
        authGroup.POST("/register", auth.RegisterHandler)
    }

    // user
    var userGroup = engine.Group(
        "/user",
        middleware.JWTMiddleware(),
    )
    {
        userGroup.POST("/change_username", user.ChangeUsernameHandler)
        userGroup.POST("/change_password", user.ChangePasswordHandler)
    }

    // file
    var fileGroup = engine.Group(
        "/file",
        middleware.JWTMiddleware(),
    )
    {
        fileGroup.GET("/storage_usage", file.StorageUsageHandler)

        fileGroup.POST("/upload", file.UploadHandler)
        fileGroup.POST("/upload_binary", file.UploadBinaryHandler)
        fileGroup.POST("/download", file.DownloadHandler)

    }

    // admin
    var adminGroup = engine.Group(
        "/admin",
        middleware.JWTMiddleware(),
    )
    {
        adminGroup.POST("/users_info", admin.UsersInfoHandler)
        adminGroup.POST("/ban_user", admin.BanUserHandler)
    }

}
