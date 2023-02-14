package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddBuyRecord(c *gin.Context) {
	// c.JSON：返回JSON格式的数据
	c.String(http.StatusOK, "123")
}
