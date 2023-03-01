package controller

import (
	"github.com/basegin/base/log"
	models "github.com/basegin/model"
	RedisUtil "github.com/basegin/redis"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ConfigParam struct {
	Key string `form:"key" binding:"required"`
}

func GetConfig(c *gin.Context) {
	var input ConfigParam
	if err := c.ShouldBind(&input); err != nil {
		log.ErrorLogger(c, err.Error())
		return
	}

	v, err := models.GetConfigValue(c, input.Key)
	if err != nil {
		log.RenderJsonFail(c, err)
	}

	log.RenderJsonSucc(c, gin.H{"value": v})
}

func SetRedis(c *gin.Context) {
	RedisUtil.RedisClient.SetStr(c.Query("key"), c.Query("val"))
	c.JSON(http.StatusOK, log.ResultSuccess)
}

func GetRedis(c *gin.Context) {
	key := c.Query("val")
	val, _ := RedisUtil.RedisClient.GetStr(key)
	c.JSON(http.StatusOK, log.ResultMsgData("设置成功", val))
}
