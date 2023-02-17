package model

import (
	"time"
	"zixishi-background/config"
)

type StudyRoom struct {
	ID    int64  `json:"id" gorm:"column:id"`
	Intro string `json:"intro"`
	Name  string `json:"name"`
}

type Seatinfo struct {
	ID            int64  `json:"id" gorm:"column:id"`
	Sid           int64  `json:"sid"`
	SeatTypeId    int64  `json:"seat_type_id"`
	CreateTime    int64  `json:"create_time"`
	UpdateTime    int64  `json:"update_time"`
	StudyRoomName string `json:"study_room_name" gorm:"-"`
	SeatTypeName  string `json:"seat_type_name" gorm:"-"`
	BuyTime       string `json:"buy_time" gorm:"-"`
}

type SeatType struct {
	ID            int64  `json:"id" gorm:"column:id"`
	Sid           int64  `json:"sid"`
	Intro         string `json:"intro"`
	PriceIntro    string `json:"price_intro"`
	Name          string `json:"name"`
	StudyRoomName string `json:"study_room_name" gorm:"-"`
}

type BuyRecord struct {
	ID        int64  `json:"id" gorm:"column:id"`
	Sid       int64  `json:"sid"`
	SeatId    int64  `json:"seat_id"`
	Uid       string `json:"uid"`
	Uname     string `json:"uname"`
	Price     int64  `json:"price"`
	EndTime   int64  `json:"end_time"`
	StartTime int64  `json:"start_time"`
	BuyTime   int64  `json:"buy_time"`
}

func GetAllStudyRoomMap() (map[int64]*StudyRoom, error) {
	studyRooms := make([]*StudyRoom, 0)
	result := config.DB.Find(&studyRooms)
	if result.Error != nil {
		return nil, result.Error
	}

	studyRoomMap := make(map[int64]*StudyRoom)
	for _, v := range studyRooms {
		studyRoomMap[v.ID] = v
	}

	return studyRoomMap, nil
}

func GetAllSeatTypeMap() (map[int64]*SeatType, error) {
	seatTypes := make([]*SeatType, 0)
	result := config.DB.Find(&seatTypes)
	if result.Error != nil {
		return nil, result.Error
	}

	seatTypeMap := make(map[int64]*SeatType)
	for _, v := range seatTypes {
		seatTypeMap[v.ID] = v
	}

	return seatTypeMap, nil
}

func GetBuyRecordMap() (map[int64][]*BuyRecord, error) {
	buyRecords := make([]*BuyRecord, 0)
	result := config.DB.Where("end_time > ?", time.Now().Unix()).Find(&buyRecords)
	if result.Error != nil {
		return nil, result.Error
	}

	buyRecordMap := make(map[int64][]*BuyRecord)
	for _, v := range buyRecords {
		if _, ok := buyRecordMap[v.SeatId]; !ok {
			buyRecordMap[v.SeatId] = make([]*BuyRecord, 0)
		}
		buyRecordMap[v.SeatId] = append(buyRecordMap[v.SeatId], v)
	}

	return buyRecordMap, nil
}
