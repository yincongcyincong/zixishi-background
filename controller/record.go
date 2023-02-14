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

	if param.BuyTime == 0 {
		param.BuyTime = time.Now().Unix()
	}

	switch param.Type {
	case 1:
		param.EndTime += 24 * 60 * 60
	case 2:
		param.EndTime += 7 * 24 * 60 * 60
	case 3:
		param.EndTime += 31 * 24 * 60 * 60
	case 4:
		param.EndTime += 92 * 24 * 60 * 60
	case 5:
		param.EndTime += 183 * 24 * 60 * 60
	case 6:
		param.EndTime += 365 * 24 * 60 * 60
	}

	// c.JSON：返回JSON格式的数据
	dbParam := &model.BuyRecord{
		Sid:     param.Sid,
		Uid:     param.Uid,
		Uname:   param.Uname,
		Price:   param.Price,
		EndTime: param.EndTime,
		BuyTime: param.BuyTime,
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
