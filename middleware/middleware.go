package middleware

import (
    "net/http"

    "github.com/gin-gonic/gin"

    "FileServerWeb/widget/jwt"
    R "FileServerWeb/widget/response"
)

// 验证请求中的 Authorization 字段, 验证成功会在header中写入 UUID 字段
// 读取: c.GetString("UUID")
func JWTMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        var err error
        var token = c.GetHeader("Authorization")
        var claims *jwt.Claims

        if token != "" {
            claims, err = jwt.ParseToken(token)
            if err != nil {
                c.AbortWithStatusJSON(
                    http.StatusUnauthorized,
                    R.Unauthorized(nil),
                )
                return
            }
        }

        c.Set("UUID", claims.UUID)
        return
    }
}


func LoggerMiddleWare() gin.HandlerFunc {
    return func(c *gin.Context) {
        // TODO
    }
}