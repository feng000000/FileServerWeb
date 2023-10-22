package main

import (
	"fmt"
	"github.com/gin-gonic/gin"

	"FileServerWeb/config"
	"FileServerWeb/routers"
	// "FileServerWeb/widget/jwt"
)

func main() {
	fmt.Println(config.HOME_DIR)
	fmt.Println(config.FILE_PATH)

	var engine = gin.Default()

	engine.Static("/static", "static")

	engine.SetTrustedProxies(nil)

	routers.Routers(engine)

	// var s string
	// s, err = jwt.GenerateToken("feng")
	// if err != nil {
	// 	fmt.Println("generate jwt failed, ", err)
	// 	return
	// }
	// fmt.Printf("token: %s\n", s)

    // // 解析jwt
	// var claim *jwt.Claims
	// claim, err = jwt.ParseToken(s)
	// if err != nil {
	// 	fmt.Println("parse jwt failed:", err)
	// 	return
	// }
	// fmt.Printf("claim: %+v\n", claim)
	// fmt.Println("username: ", claim.Username)

	engine.Run("127.0.0.1:8080")
}
