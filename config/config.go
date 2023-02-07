package config

import "github.com/medivhzhan/weapp/v3"

var SDK *weapp.Client

func InitConfig() {
	SDK = weapp.NewClient("your-appid", "your-secret")
}
