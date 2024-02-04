package main

import (
    "time"

    "github.com/gin-gonic/gin"
    ginzap "github.com/gin-contrib/zap"

    "FileServerWeb/routers"
    "FileServerWeb/widget/logger"
    "FileServerWeb/config"
)


func main() {
    defer logger.Logger.Sync()

    if config.DEBUG {
        gin.SetMode(gin.DebugMode)
    } else {
        gin.SetMode(gin.ReleaseMode)
    }

    var engine = gin.New()

    engine.Static("/static", "static")

    // 使用日志中间件
    engine.Use(ginzap.Ginzap(logger.Logger, time.RFC3339, false))
    engine.Use(ginzap.RecoveryWithZap(logger.Logger, true))

    engine.SetTrustedProxies(nil)

    routers.Routers(engine)

    engine.Run("127.0.0.1:8080")
}
