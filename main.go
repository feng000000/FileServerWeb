package main

import (
	"github.com/gin-gonic/gin"

	"FileServerWeb/routers"
	"FileServerWeb/widget/logger"
)

func main() {
	defer logger.Logger.Sync()

	var engine = gin.Default()

	engine.Static("/static", "static")

	engine.SetTrustedProxies(nil)

	routers.Routers(engine)

	engine.Run("127.0.0.1:8080")

}
