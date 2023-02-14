package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"zixishi-background/config"
	"zixishi-background/model"
	"zixishi-background/utils"
)

func GetStudyRoom(c *gin.Context) {
	page := c.GetInt(c.DefaultQuery("page", "1"))

	number := 20
	limit := (page - 1) * number

	records := make([]*model.StudyRoom, 0)
	result := config.DB.Limit(limit).Order("id desc").Find(&records)
	if result.Error != nil {
		log.Println(result.Error)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	var count int64
	result = config.DB.Count(&count)
	if result.Error != nil {
		log.Println(result.Error)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	// c.JSON：返回JSON格式的数据
	c.HTML(http.StatusOK, "main.html", gin.H{
		"title":   "posts/index",
		"records": records,
		"count":   utils.NewPaginator(c.Request, number, count),
	})

}

func AddStudyRoom(c *gin.Context) {
	param := new(model.StudyRoomParam)
	err := c.BindJSON(param)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, utils.Fail(utils.ParamErrCode, utils.ParamErrMsg, ""))
		return
	}

	// c.JSON：返回JSON格式的数据
	dbParam := &model.StudyRoom{
		Name:  param.Name,
		Intro: param.Intro,
	}
	result := config.DB.Create(dbParam)
	if result.Error != nil {
		log.Println(result.Error)
		c.JSON(http.StatusOK, utils.Fail(utils.DBErrCode, utils.DBErrMsg, ""))
		return
	}

	c.JSON(http.StatusOK, utils.Succ(""))

}

func UpdateStudyRoom(c *gin.Context) {
	param := new(model.StudyRoomParam)
	err := c.BindJSON(param)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, utils.Fail(utils.ParamErrCode, utils.ParamErrMsg, ""))
		return
	}

	// c.JSON：返回JSON格式的数据
	dbParam := &model.StudyRoom{
		ID:    param.ID,
		Name:  param.Name,
		Intro: param.Intro,
	}
	result := config.DB.Save(dbParam)
	if result.Error != nil {
		log.Println(result.Error)
		c.JSON(http.StatusOK, utils.Fail(utils.DBErrCode, utils.DBErrMsg, ""))
		return
	}

	c.JSON(http.StatusOK, utils.Succ(""))

}

func DeleteStudyRoom(c *gin.Context) {
	param := new(model.StudyRoomParam)
	err := c.BindJSON(param)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, utils.Fail(utils.ParamErrCode, utils.ParamErrMsg, ""))
		return
	}

	// c.JSON：返回JSON格式的数据
	dbParam := &model.StudyRoom{
		ID: param.ID,
	}
	result := config.DB.Delete(dbParam)
	if result.Error != nil {
		log.Println(result.Error)
		c.JSON(http.StatusOK, utils.Fail(utils.DBErrCode, utils.DBErrMsg, ""))
		return
	}

	c.JSON(http.StatusOK, utils.Succ(""))

}
