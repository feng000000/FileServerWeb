package auth

import (
	// "log"
	// "fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"FileServerWeb/db"
	"FileServerWeb/widget/jwt"
	r "FileServerWeb/widget/response"
)

var DB = db.DB

var query_user, err = DB.Preparex(`select * from users where username=? and password=?`)

func Login(c *gin.Context) {
	var err error
	var username string
	var password string
	var token = c.GetHeader("Authorization")

	if token != "" {
		_, err = jwt.ParseToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, r.StatusUnauthorized(r.Json{
				"message": "Invalid token",
			}))
			return
		}
		c.JSON(http.StatusOK, r.Success(nil))
		return
	}

	var json = make(map[string]interface{})
	c.BindJSON(&json)

	username = json["username"].(string)
	password = json["password"].(string)

	var user db.UsersTable
	err = query_user.Get(&user, username, password)

	if err != nil {
		c.JSON(http.StatusUnauthorized, r.StatusUnauthorized(r.Json{
			"message": "Wrong username or password",
		}))
		return
	}

	var ret_token string
	ret_token, err = jwt.GenerateToken(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "InternalServerError",
		})
		return
	}
	c.JSON(http.StatusOK, r.Success(r.Json{
		"token" : ret_token,
	}))
	return
}
