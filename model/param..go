package model

type StudyRoomParam struct {
	ID    int64  `json:"id"`
	Intro string `json:"intro"`
	Name  string `json:"name"`
}

type SeatinfoParam struct {
	ID         int64 `json:"id"`
	Sid        int64 `json:"sid"`
	SeatTypeId int64 `json:"seat_type_id"`
}

type SeatTypeParam struct {
	ID         int64  `json:"id"`
	Sid        int64  `json:"sid"`
	Intro      string `json:"intro"`
	PriceIntro string `json:"price_intro"`
	Name       string `json:"name"`
}

type BuyRecordParam struct {
	Sid       int64  `json:"sid"`
	SeatId    int64  `json:"seat_id"`
	Uid       string `json:"uid"`
	Uname     string `json:"uname"`
	Price     int64  `json:"price"`
	EndTime   int64  `json:"end_time"`
	StartTime int64  `json:"start_time"`
	BuyTime   int64  `json:"buy_time"`
}
