package main

import (
	"github.com/gin-gonic/gin"
	"zixishi-background/controller"
)

func main() {

	r := gin.Default()
	r.GET("/ping", controller.Ping)
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}

