package main

import (
	"fmt"
	"github.com/gin-gonic/gin"

	"FileServerWeb/config"
	"FileServerWeb/routers"
	"FileServerWeb/widget/jwt"
)

func main() {
	fmt.Println(config.HOME_DIR)
	fmt.Println(config.FILE_PATH)

	r := gin.Default()

	r.Static("/static", "static")

	routers.Routers(r)

	s, err := jwt.GenerateToken("feng")
	if err != nil {
		fmt.Println("generate jwt failed, ", err)
		return
	}
	fmt.Printf("token: %s\n", s)

    // 解析jwt
	claim, err := jwt.ParseToken(s)
	if err != nil {
		fmt.Println("parse jwt failed:", err)
		return
	}
	fmt.Printf("claim: %+v\n", claim)
	fmt.Println("username: ", claim.Username)

	r.Run()
}
