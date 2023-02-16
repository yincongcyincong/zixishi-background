package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/medivhzhan/weapp/v3"
	"github.com/unknwon/goconfig"
)

var (
	SDK  *weapp.Client
	DB   *gorm.DB
	Conf *Config
)

type Config struct {
	AppId     string
	AppSecret string

	DBConf string
}

func init() {
	initConfig()
	initWxInfo()
	initDb()
}

func initConfig() {
	C, err := goconfig.LoadConfigFile("./config/conf.ini")
	if err != nil {
		panic(err)
	}
	Conf.AppId = C.MustValue("wx", "appId", "")
	Conf.AppSecret = C.MustValue("wx", "appSecret", "")
	Conf.DBConf = C.MustValue("wx", "DBConf", "")
}

func initWxInfo() {
	tokenGetter := func(appid, secret string) (token string, expireIn uint) {
		expireIn = 1000
		token = "your-custom-token"
		return token, expireIn
	}

	SDK = weapp.NewClient(
		Conf.AppId,
		Conf.AppSecret,
		weapp.WithAccessTokenSetter(tokenGetter),
	)
}

func initDb() {
	var err error
	DB, err = gorm.Open("mysql", Conf.DBConf)
	if err != nil {
		panic("failed to connect database")
	}
}
