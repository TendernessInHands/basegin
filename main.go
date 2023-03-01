package main

import (
	"fmt"
	"github.com/basegin/base"
	"github.com/basegin/db"
	RedisUtil "github.com/basegin/redis"
	"github.com/basegin/router"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	// 初始化 gin 框架
	r := gin.Default()
	fmt.Print("r", r)

	//启动配置
	base.Bootstrap(r)
	//初始化mysql
	db.InitMysql()
	//初始化redis
	RedisUtil.InitRedis()
	// 添加路由
	router.Http(r)

	r.GET("/health/:val", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "请求成功!",
			"data": c.Param("val"),
		})
	})

	//启动监听端口
	endless.ListenAndServe(":8080", r)
}
