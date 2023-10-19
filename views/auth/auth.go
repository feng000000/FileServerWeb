package auth

import (
	// "fmt"
	"FileServerWeb/widget/jwt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	token := c.GetHeader("Authorization")
	res, err := jwt.ParseToken(token)
	username := (*res).Username
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Not login",
		})
		return
	}
	json := make(map[string]interface{})
	c.BindJSON(&json)

	// password := json["password"]

	ret_token, err := jwt.GenerateToken(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "InternalServerError",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message" : "Success",
		"token" : ret_token,
	})
}