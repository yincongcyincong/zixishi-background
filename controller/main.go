package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Main(c *gin.Context) {
	// c.JSON：返回JSON格式的数据
	c.HTML(http.StatusOK, "main.html", gin.H{
		"title": "posts/index",
	})
}