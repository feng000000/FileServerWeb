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
    var ip string
    ip = c.Request.Header.Get("X-Real-IP")
    if ip == "127.0.0.1" || ip == "" {
        ip = c.Request.Header.Get("X-Forwarded-For")
    }
    if ip == "" {
        ip = "127.0.0.1"
    }

    c.JSON(http.StatusOK, gin.H{
        "ip": ip,
    })
}
