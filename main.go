package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"zixishi-background/config"
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
	r.Run(config.Conf.Port)
}

func initRouter(r *gin.Engine) {
	r.LoadHTMLFiles("views/main.html", "views/seatinfo.html", "views/paginator.html", "views/login.html",
		"views/seatinfo_form.html", "views/studyroom_form.html", "views/studyroom.html", "views/seat_type.html",
		"views/seat_type_form.html", "views/buy_record.html", "views/buy_record_form.html")
	r.StaticFS("/static/style", http.Dir("./static/style"))

	r.Use(middleware.AuthMiddleWare())

	r.GET("/main", controller.Main)
	r.GET("/main/default", controller.GetSeatInfo)

	r.GET("/study_room/get", controller.GetStudyRoom)
	r.POST("/study_room/add", controller.AddStudyRoom)
	r.GET("/study_room/form", controller.StudyRoomForm)
	r.POST("/study_room/update", controller.UpdateStudyRoom)
	r.POST("/study_room/delete", controller.DeleteStudyRoom)

	r.GET("/seatinfo/get", controller.GetSeatInfo)
	r.POST("/seatinfo/add", controller.AddSeatInfo)
	r.GET("/seatinfo/form", controller.SeatInfoForm)
	r.POST("/seatinfo/update", controller.UpdateSeatInfo)
	r.POST("/seatinfo/delete", controller.DeleteSeatInfo)

	r.GET("/seat_type/get", controller.GetSeatType)
	r.POST("/seat_type/add", controller.AddSeatType)
	r.GET("/seat_type/form", controller.SeatTypeForm)
	r.POST("/seat_type/update", controller.UpdateSeatType)
	r.POST("/seat_type/delete", controller.DeleteSeatType)

	r.GET("/buy_record/get", controller.GetRecord)
	r.GET("/buy_record/form", controller.BuyRecordForm)
	r.POST("/buy_record/add", controller.AddBuyRecord)
	r.POST("/buy_record/update", controller.UpdateBuyRecord)

	r.GET("/login", controller.Login)
	r.POST("/login", controller.DoLogin)
	r.GET("/login/captcha", controller.Captcha)

	r.POST("/img/upload", controller.Upload)

}
