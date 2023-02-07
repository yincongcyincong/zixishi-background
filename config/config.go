package config

import "github.com/medivhzhan/weapp/v3"

var SDK *weapp.Client

func InitConfig() {
	SDK = weapp.NewClient("your-appid", "your-secret")

	tokenGetter := func(appid, secret string) (token string, expireIn uint) {
		expireIn = 1000
		token = "your-custom-token"
		return token, expireIn
	}

	sdk := weapp.NewClient(
		"your-appid",
		"your-secret",
		weapp.WithAccessTokenSetter(tokenGetter),
	)
}
