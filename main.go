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

	engine.Run("127.0.0.1:8080")
}
