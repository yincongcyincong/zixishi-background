package model

type StudyRoom struct {
	ID    int64  `json:"id" gorm:"column:id"`
	Intro string `json:"intro"`
	Name  string `json:"name"`
}

type Seatinfo struct {
	ID         int64 `json:"id" gorm:"column:id"`
	Sid        int64 `json:"sid"`
	SeatTypeId int64 `json:"seat_type_id"`
	CreateTime int64 `json:"create_time"`
	UpdateTime int64 `json:"update_time"`
}

type SeatType struct {
	ID         int64  `json:"id" gorm:"column:id"`
	Sid        int64  `json:"sid"`
	Intro      string `json:"intro"`
	PriceIntro string `json:"price_intro"`
	Name       string `json:"name"`
}

type BuyRecord struct {
	ID      int64  `json:"id" gorm:"column:id"`
	Sid     int64  `json:"sid"`
	Uid     string `json:"uid"`
	Uname   string `json:"uname"`
	Price   int64  `json:"price"`
	EndTime int64  `json:"end_time"`
	BuyTime int64  `json:"buy_time"`
}
