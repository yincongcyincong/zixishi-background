package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"zixishi-background/config"
	"zixishi-background/model"
	"zixishi-background/utils"
)

func DoLogin(c *gin.Context) {
	param := new(model.LoginParam)
	err := c.BindJSON(param)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, utils.Fail(utils.ParamErrCode, utils.ParamErrMsg, ""))
		return
	}

	if param.Username != config.Conf.UserName && param.Password != config.Conf.Password {
		log.Println("user or password unmatch")
		c.JSON(http.StatusOK, utils.Fail(utils.PassErrCode, utils.PassErrMsg, ""))
		return
	}

	// c.JSON：返回JSON格式的数据
	cookie := &http.Cookie{
		Name:     "session_id",
		Value:    "session-886", //这个是value要自己生成？？规则自定就可以？
		Path:     "/",
		HttpOnly: true,
	}
	http.SetCookie(c.Writer, cookie)
	c.JSON(http.StatusOK, utils.Succ(""))
}

func Login(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{
		"title":   "登录",
	})
}
