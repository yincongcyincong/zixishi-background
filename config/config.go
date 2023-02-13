package config

import "github.com/medivhzhan/weapp/v3"

var SDK *weapp.Client

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
