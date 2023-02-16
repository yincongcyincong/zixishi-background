package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"zixishi-background/controller"
	"zixishi-background/middleware"
	"zixishi-background/utils"
)

func main() {
	r := gin.Default()
	r.SetFuncMap(template.FuncMap{
		"dateFormat": utils.NewDate().Format,
	})
	initRouter(r)
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}

func initRouter(r *gin.Engine) {
	r.LoadHTMLFiles("views/main.html", "views/seatinfo.html", "views/paginator.html", "views/login.html")
	r.StaticFS("/static", http.Dir("./static"))

	r.Use(middleware.AuthMiddleWare())

	r.GET("/ping", controller.Ping)
	r.GET("/main/default", controller.GetSeatInfo)

	r.GET("/study_room/get", controller.GetStudyRoom)
	r.POST("/study_room/add", controller.AddStudyRoom)
	r.POST("/study_room/update", controller.UpdateStudyRoom)
	r.POST("/study_room/del", controller.DeleteStudyRoom)

	r.GET("/seatinfo/get", controller.GetSeatInfo)
	r.POST("/seatinfo/add", controller.AddSeatInfo)
	r.POST("/seatinfo/update", controller.UpdateSeatInfo)
	r.POST("/seatinfo/del", controller.DeleteSeatInfo)

	r.GET("/seat_type/get", controller.GetSeatType)
	r.POST("/seat_type/add", controller.AddSeatType)
	r.POST("/seat_type/update", controller.UpdateSeatType)
	r.POST("/seat_type/del", controller.DeleteSeatType)

	r.GET("/buy_record/get", controller.GetRecord)
	r.POST("/buy_record/add", controller.AddBuyRecord)

	r.GET("/login", controller.Login)
	r.GET("/do_login", controller.DoLogin)

}
