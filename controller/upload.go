package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"path"
)

func Upload(c *gin.Context) {
	file, err := c.FormFile("wangeditor-uploaded-image")
	if err != nil {
		log.Println("controller - admin - upload", err)
		c.JSON(http.StatusOK, gin.H{
			"errno":   1,
			"message": "read body fail",
		})
		return
	}

	filepath := path.Join("./static/upload/", file.Filename)
	err = c.SaveUploadedFile(file, filepath)
	if err != nil {
		log.Println("controller - admin - upload - SaveUploadedFile", err)
		c.JSON(http.StatusOK, map[string]interface{}{
			"errno":   1,
			"message": "save file fail",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"errno":   0,
		"message": "success",
		"data": map[string]interface{}{
			"url": "http://" + c.Request.Host + "/static/upload/" + file.Filename,
		},
	})
}
