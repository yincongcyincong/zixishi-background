package controller

import (
	"encoding/json"
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

	studyRoomMap, err := model.GetAllStudyRoomMap()
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
		"title":     "座位信息",
		"records":   records,
		"paginator": utils.NewPaginator(c.Request, number, count),
	})

}

func SeatInfoForm(c *gin.Context) {
	id := c.Query("id")
	seatinfo := new(model.Seatinfo)
	config.DB.Where("id = ?", id).First(seatinfo)

	seatTypeMap, err := model.GetAllSeatTypeMap()
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	studyRoomMap, err := model.GetAllStudyRoomMap()
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	// c.JSON：返回JSON格式的数据
	c.HTML(http.StatusOK, "seatinfo_form.html", gin.H{
		"title":        "座位信息",
		"seatinfo":     seatinfo,
		"seatTypeMap":  seatTypeMap,
		"studyRoomMap": studyRoomMap,
	})

}

func AddSeatInfo(c *gin.Context) {
	param := new(model.SeatinfoParam)
	err := c.Bind(param)
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

	c.JSON(http.StatusOK, utils.SuccWithUrl("/seatinfo/get"))

}

func UpdateSeatInfo(c *gin.Context) {
	param := new(model.SeatinfoParam)
	err := c.Bind(param)
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

	c.JSON(http.StatusOK, utils.SuccWithUrl("/seatinfo/get"))

}

func DeleteSeatInfo(c *gin.Context) {
	param := new(model.SeatinfoParam)
	err := c.Bind(param)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, utils.Fail(utils.ParamErrCode, utils.ParamErrMsg, ""))
		return
	}

	if param.ID == 0 {
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

	c.JSON(http.StatusOK, utils.SuccWithUrl("/seatinfo/get"))

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
	result = config.DB.Model(records).Count(&count)
	if result.Error != nil {
		log.Println(result.Error)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	studyRoomMap, err := model.GetAllStudyRoomMap()
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	for _, v := range records {
		if _, ok := studyRoomMap[v.Sid]; ok {
			v.StudyRoomName = studyRoomMap[v.Sid].Name
		}
		json.Unmarshal([]byte(v.PriceIntro), &v.PriceInfo)
	}

	// c.JSON：返回JSON格式的数据
	c.HTML(http.StatusOK, "seat_type.html", gin.H{
		"title":     "seat_type",
		"records":   records,
		"paginator": utils.NewPaginator(c.Request, number, count),
	})

}

func AddSeatType(c *gin.Context) {
	param := new(model.SeatTypeParam)
	err := c.Bind(param)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, utils.Fail(utils.ParamErrCode, utils.ParamErrMsg, ""))
		return
	}

	fmt.Println(fmt.Sprintf("%+v", param))

	priceInfo := make(map[string]int64)
	for idx, val := range param.Key {
		priceInfo[val] = param.Price[idx]
	}

	tmp, _ := json.Marshal(priceInfo)
	param.PriceIntro = string(tmp)

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

	c.JSON(http.StatusOK, utils.SuccWithUrl("/seat_type/get"))

}

func SeatTypeForm(c *gin.Context) {
	id := c.Query("id")
	seatType := new(model.SeatType)
	config.DB.Where("id = ?", id).First(seatType)

	if seatType.PriceIntro != "" {
		json.Unmarshal([]byte(seatType.PriceIntro), &seatType.PriceInfo)
	}

	studyRoomMap, err := model.GetAllStudyRoomMap()
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	// c.JSON：返回JSON格式的数据
	c.HTML(http.StatusOK, "seat_type_form.html", gin.H{
		"title":        "座位类型信息",
		"seatType":     seatType,
		"studyRoomMap": studyRoomMap,
	})

}

func UpdateSeatType(c *gin.Context) {
	param := new(model.SeatTypeParam)
	err := c.Bind(param)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, utils.Fail(utils.ParamErrCode, utils.ParamErrMsg, ""))
		return
	}

	priceInfo := make(map[string]int64)
	for idx, val := range param.Key {
		priceInfo[val] = param.Price[idx]
	}

	tmp, _ := json.Marshal(priceInfo)
	param.PriceIntro = string(tmp)

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

	c.JSON(http.StatusOK, utils.SuccWithUrl("/seat_type/get"))

}

func DeleteSeatType(c *gin.Context) {
	param := new(model.SeatTypeParam)
	err := c.Bind(param)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, utils.Fail(utils.ParamErrCode, utils.ParamErrMsg, ""))
		return
	}

	if param.ID == 0 {
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

	c.JSON(http.StatusOK, utils.SuccWithUrl("/seat_type/get"))

}
