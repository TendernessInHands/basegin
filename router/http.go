package router

import (
	"github.com/basegin/controller"
	"github.com/gin-gonic/gin"
)

func Http(g *gin.Engine) {

	api := g.Group("/api")

	api.POST("/getConfig", controller.GetConfig)

	api.GET("/setRedis", controller.SetRedis)

	api.GET("/getRedis", controller.GetRedis)
}
