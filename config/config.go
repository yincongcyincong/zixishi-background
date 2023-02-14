package config

import (
	"github.com/jinzhu/gorm"
	"github.com/medivhzhan/weapp/v3"
)

var (
	SDK *weapp.Client
	DB  *gorm.DB
)

func init() {
	InitConfig()
	InitDb()
}

func InitConfig() {

	tokenGetter := func(appid, secret string) (token string, expireIn uint) {
		expireIn = 1000
		token = "your-custom-token"
		return token, expireIn
	}

	SDK = weapp.NewClient(
		"your-appid",
		"your-secret",
		weapp.WithAccessTokenSetter(tokenGetter),
	)
}

func InitDb() {
	var err error
	DB, err = gorm.Open("mysql", "root:12345@/demo?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
}
