package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	// c.JSON：返回JSON格式的数据
	cookie := &http.Cookie{
		Name:     "session_id",
		Value:    "onion", //这个是value要自己生成？？规则自定就可以？
		Path:     "/",
		HttpOnly: true,
	}
	http.SetCookie(c.Writer, cookie)
	c.String(http.StatusOK, "登录成功")
}
