package model

type StudyRoomParam struct {
	ID    int64  `json:"id" form:"id"`
	Intro string `json:"intro" form:"intro"`
	Name  string `json:"name" form:"name"`
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
	ID        int64  `json:"id"  form:"id"`
	Sid       int64  `json:"sid"  form:"sid"`
	SeatId    int64  `json:"seat_id"  form:"seat_id"`
	Uid       string `json:"uid"`
	Uname     string `json:"uname" form:"uname"`
	Price     int64  `json:"price" form:"price"`
	EndTime   int64  `json:"end_time" form:"end_time"`
	StartTime int64  `json:"start_time" form:"start_time"`
	BuyTime   int64  `json:"buy_time"`
}

type LoginParam struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
