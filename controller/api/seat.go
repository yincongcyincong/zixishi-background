package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"sort"
	"time"
	"zixishi-background/model"
	"zixishi-background/utils"
)

func GetSeatInfo(c *gin.Context) {
	// c.JSON：返回JSON格式的数据
	seatinfos, err := model.GetAllSeatinfo()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, utils.Fail(utils.DBErrCode, utils.DBErrMsg, ""))
		return
	}

	seatTypeMap, err := model.GetAllSeatTypeMap()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, utils.Fail(utils.DBErrCode, utils.DBErrMsg, ""))
		return
	}

	seatTypes := make([]*model.SeatType, 0)
	for _, tmp := range seatTypeMap {
		seatTypes = append(seatTypes, tmp)
	}

	sort.Slice(seatTypes, func(i, j int) bool {
		return seatTypes[i].ID < seatTypes[j].ID
	})

	buyRecordMap, err := model.GetBuyRecordMap()
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	for _, v := range seatinfos {
		if _, ok := seatTypeMap[v.SeatTypeId]; ok {
			v.SeatTypeName = seatTypeMap[v.SeatTypeId].Name
		}
		if _, ok := buyRecordMap[v.SeatTypeId]; ok {
			for _, r := range buyRecordMap[v.SeatTypeId] {
				v.BuyTime = fmt.Sprintf("%s-%s,", time.Unix(r.StartTime, 0).Format("2006-01-02"), time.Unix(r.EndTime, 0).Format("2006-01-02"))
			}
		}
	}

	c.JSON(http.StatusOK, utils.Succ(map[string]interface{}{
		"seatType": seatTypes,
		"seatinfo": seatinfos,
	}))
}
