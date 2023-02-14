package model

type StudyRoomParam struct {
	Intro string `json:"intro"`
	Name  string `json:"name"`
}

type SeatinfoParam struct {
	Sid        int64 `json:"sid"`
	SeatTypeId int64 `json:"seat_type_id"`
	UpdateTime int64 `json:"update_time"`
}

type SeatTypeParam struct {
	Sid        int64  `json:"sid"`
	Intro      string `json:"intro"`
	PriceIntro string `json:"price_intro"`
	Name       string `json:"name"`
}

type BuyRecordParam struct {
	Sid     int64  `json:"sid"`
	Uid     string `json:"uid"`
	Uname   string `json:"uname"`
	Price   int64  `json:"price"`
	EndTime int64  `json:"end_time"`
	BuyTime int64  `json:"buy_time"`
	Type    int    `json:"type"` // 1 一天 2 一周 3 一月 4 三个月 5 6个月 6 一年
}
