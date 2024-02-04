package middleware

import (
    "net/http"

    "github.com/gin-gonic/gin"

    "FileServerWeb/widget/auth"
    R "FileServerWeb/widget/response"
)

// 验证请求中的 Authorization 字段, 验证成功会在header中写入 UUID 字段
// 读取: c.GetString("UUID")
func JWTMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        var err error
        var token = c.GetHeader("Authorization")
        var claims *auth.Claims

        if token == "" {
            c.AbortWithStatusJSON(
                http.StatusUnauthorized,
                R.Unauthorized(nil),
            )
            return
        }

        claims, err = auth.ParseToken(token)
        if err != nil {
            c.AbortWithStatusJSON(
                http.StatusUnauthorized,
                R.Unauthorized(nil),
            )
            return
        }

        c.Set("UUID", claims.UUID)
        return
    }
}