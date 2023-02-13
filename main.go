package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"zixishi-background/controller"
	"zixishi-background/utils"
)

func main() {

	r := gin.Default()
	r.SetFuncMap(template.FuncMap{
		"dateFormat": utils.NewDate().Format,
	})
	r.LoadHTMLFiles("views/main.html")
	r.StaticFS("/static", http.Dir("./static"))
	r.GET("/ping", controller.Ping)
	r.Any("/main/default", controller.DEFAULT)
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}

