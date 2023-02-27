package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
	"zixishi-background/config"
	"zixishi-background/model"
	"zixishi-background/utils"
)

func AddBuyRecord(c *gin.Context) {

	param := new(model.BuyRecordParam)
	err := c.BindJSON(param)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, utils.Fail(utils.ParamErrCode, utils.ParamErrMsg, ""))
		return
	}

	if param.Sid == 0 || param.Price == 0 {
		log.Println(err)
		c.JSON(http.StatusOK, utils.Fail(utils.ParamErrCode, utils.ParamErrMsg, ""))
		return
	}

	if param.StartTime == 0 {
		param.StartTime = time.Now().Unix()
	}

	// c.JSON：返回JSON格式的数据
	dbParam := &model.BuyRecord{
		Sid:       param.Sid,
		Uid:       param.Uid,
		Uname:     param.Uname,
		Price:     param.Price,
		EndTime:   param.EndTime,
		StartTime: param.StartTime,
		BuyTime:   time.Now().Unix(),
	}
	result := config.DB.Create(dbParam)
	if result.Error != nil {
		log.Println(result.Error)
		c.JSON(http.StatusOK, utils.Fail(utils.DBErrCode, utils.DBErrMsg, ""))
		return
	}

	c.JSON(http.StatusOK, utils.Succ(""))
}

func GetRecord(c *gin.Context) {
	page := c.GetInt(c.DefaultQuery("page", "1"))

	number := 20
	limit := (page - 1) * number

	records := make([]*model.BuyRecord, 0)
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

func UpdateBuyRecord(c *gin.Context) {
	param := new(model.BuyRecordParam)
	err := c.BindJSON(param)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, utils.Fail(utils.ParamErrCode, utils.ParamErrMsg, ""))
		return
	}

	if param.StartTime == 0 {
		param.StartTime = time.Now().Unix()
	}

	// c.JSON：返回JSON格式的数据
	dbParam := &model.BuyRecord{
		ID:        param.ID,
		Sid:       param.Sid,
		EndTime:   param.EndTime,
		StartTime: param.StartTime,
	}
	result := config.DB.Model(dbParam).Update(dbParam)
	if result.Error != nil {
		log.Println(result.Error)
		c.JSON(http.StatusOK, utils.Fail(utils.DBErrCode, utils.DBErrMsg, ""))
		return
	}

	c.JSON(http.StatusOK, utils.Succ(""))
}

func BuyRecordForm(c *gin.Context) {
	id := c.Query("id")
	buyRecord := new(model.BuyRecord)
	config.DB.Where("id = ?", id).First(buyRecord)

	// c.JSON：返回JSON格式的数据
	c.HTML(http.StatusOK, "buy_record.html", gin.H{
		"title":     "购买记录",
		"buyRecord": buyRecord,
	})

}
