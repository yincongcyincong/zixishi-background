package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
	"zixishi-background/config"
	"zixishi-background/model"
	"zixishi-background/utils"
)

func GetSeatInfo(c *gin.Context) {
	page := c.GetInt(c.DefaultQuery("page", "1"))

	number := 20
	limit := (page - 1) * number

	records := make([]*model.Seatinfo, 0)
	result := config.DB.Limit(limit).Order("id desc").Find(&records)
	if result.Error != nil {
		log.Println(result.Error)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	var count int64
	result = config.DB.Model(records).Count(&count)
	if result.Error != nil {
		log.Println(result.Error)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	seatTypeMap, err := model.GetAllSeatTypeMap()
	if err != nil {
		log.Println(result.Error)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	studyRoomMap, err := model.GetAllSeatTypeMap()
	if err != nil {
		log.Println(result.Error)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	buyRecordMap, err := model.GetBuyRecordMap()
	if err != nil {
		log.Println(result.Error)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	for _, v := range records {
		if _, ok := studyRoomMap[v.Sid]; ok {
			v.StudyRoomName = studyRoomMap[v.Sid].Name
		}
		if _, ok := seatTypeMap[v.SeatTypeId]; ok {
			v.SeatTypeName = seatTypeMap[v.SeatTypeId].Name
		}
		if _, ok := buyRecordMap[v.SeatTypeId]; ok {
			for _, r := range buyRecordMap[v.SeatTypeId] {
				v.BuyTime = fmt.Sprintf("%s-%s,", time.Unix(r.StartTime, 0).Format("2006-01-02"), time.Unix(r.EndTime, 0).Format("2006-01-02"))
			}
		}
	}

	// c.JSON：返回JSON格式的数据
	c.HTML(http.StatusOK, "seatinfo.html", gin.H{
		"title":   "座位信息",
		"records": records,
		"count":   utils.NewPaginator(c.Request, number, count),
	})

}

func AddSeatInfo(c *gin.Context) {
	param := new(model.SeatinfoParam)
	err := c.BindJSON(param)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, utils.Fail(utils.ParamErrCode, utils.ParamErrMsg, ""))
		return
	}

	// c.JSON：返回JSON格式的数据
	dbParam := &model.Seatinfo{
		Sid:        param.Sid,
		SeatTypeId: param.SeatTypeId,
		CreateTime: time.Now().Unix(),
		UpdateTime: time.Now().Unix(),
	}
	result := config.DB.Create(dbParam)
	if result.Error != nil {
		log.Println(result.Error)
		c.JSON(http.StatusOK, utils.Fail(utils.DBErrCode, utils.DBErrMsg, ""))
		return
	}

	c.JSON(http.StatusOK, utils.Succ(""))

}

func UpdateSeatInfo(c *gin.Context) {
	param := new(model.SeatinfoParam)
	err := c.BindJSON(param)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, utils.Fail(utils.ParamErrCode, utils.ParamErrMsg, ""))
		return
	}

	// c.JSON：返回JSON格式的数据
	dbParam := &model.Seatinfo{
		ID:         param.ID,
		SeatTypeId: param.SeatTypeId,
		UpdateTime: time.Now().Unix(),
	}
	result := config.DB.Model(dbParam).Update(dbParam)
	if result.Error != nil {
		log.Println(result.Error)
		c.JSON(http.StatusOK, utils.Fail(utils.DBErrCode, utils.DBErrMsg, ""))
		return
	}

	c.JSON(http.StatusOK, utils.Succ(""))

}

func DeleteSeatInfo(c *gin.Context) {
	param := new(model.SeatinfoParam)
	err := c.BindJSON(param)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, utils.Fail(utils.ParamErrCode, utils.ParamErrMsg, ""))
		return
	}

	// c.JSON：返回JSON格式的数据
	dbParam := &model.Seatinfo{
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

func GetSeatType(c *gin.Context) {
	page := c.GetInt(c.DefaultQuery("page", "1"))

	number := 20
	limit := (page - 1) * number

	records := make([]*model.SeatType, 0)
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

func AddSeatType(c *gin.Context) {
	param := new(model.SeatTypeParam)
	err := c.BindJSON(param)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, utils.Fail(utils.ParamErrCode, utils.ParamErrMsg, ""))
		return
	}

	// c.JSON：返回JSON格式的数据
	dbParam := &model.SeatType{
		Sid:        param.Sid,
		Intro:      param.Intro,
		PriceIntro: param.PriceIntro,
		Name:       param.Name,
	}
	result := config.DB.Create(dbParam)
	if result.Error != nil {
		log.Println(result.Error)
		c.JSON(http.StatusOK, utils.Fail(utils.DBErrCode, utils.DBErrMsg, ""))
		return
	}

	c.JSON(http.StatusOK, utils.Succ(""))

}

func UpdateSeatType(c *gin.Context) {
	param := new(model.SeatTypeParam)
	err := c.BindJSON(param)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, utils.Fail(utils.ParamErrCode, utils.ParamErrMsg, ""))
		return
	}

	// c.JSON：返回JSON格式的数据
	dbParam := &model.SeatType{
		ID:         param.ID,
		Intro:      param.Intro,
		PriceIntro: param.PriceIntro,
		Name:       param.Name,
	}
	result := config.DB.Model(dbParam).Update(dbParam)
	if result.Error != nil {
		log.Println(result.Error)
		c.JSON(http.StatusOK, utils.Fail(utils.DBErrCode, utils.DBErrMsg, ""))
		return
	}

	c.JSON(http.StatusOK, utils.Succ(""))

}

func DeleteSeatType(c *gin.Context) {
	param := new(model.SeatTypeParam)
	err := c.BindJSON(param)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, utils.Fail(utils.ParamErrCode, utils.ParamErrMsg, ""))
		return
	}

	// c.JSON：返回JSON格式的数据
	dbParam := &model.SeatType{
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
