package routers

import (
    // "fmt"
    "github.com/gin-gonic/gin"

    "FileServerWeb/middleware"
    "FileServerWeb/views"
    "FileServerWeb/views/admin"
    "FileServerWeb/views/authentication"
    "FileServerWeb/views/file"
    "FileServerWeb/views/userSettings"
)

func Routers(engine *gin.Engine) {
    engine.GET("/ping", views.Ping)
    engine.GET("/ip", views.IP)

    // auth
    var authGroup = engine.Group("/auth")
    {
        authGroup.POST("/login", authentication.LoginHandler)
        authGroup.POST("/register", authentication.RegisterHandler)
    }

    // user
    var userGroup = engine.Group(
        "/user_settings",
        middleware.JWTMiddleware(),
    )
    {
        userGroup.POST("/change_username",
                       userSettings.ChangeUsernameHandler)
        userGroup.POST("/change_password",
                       userSettings.ChangePasswordHandler)
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
