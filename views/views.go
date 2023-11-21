package views

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": "pong",
    })
}

func IP(c *gin.Context) {
    // TODO: 配置nginx后, 获取 headers 中 X-Real-IP 字段的值
    c.JSON(http.StatusOK, gin.H{
        "ip": c.ClientIP(),
    })
}
