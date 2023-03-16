package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"zixishi-background/config"
	"zixishi-background/model"
	"zixishi-background/utils"
)

func GetStudyRoom(c *gin.Context) {
	var studyRoom *model.StudyRoom
	id := c.Query("id")
	if id != "" {
		studyRoom = new(model.StudyRoom)
		config.DB.Where("id = ?", id).First(studyRoom)
	}

	c.JSON(http.StatusOK, utils.Succ(map[string]interface{}{
		"studyRoom": studyRoom,
	}))
}
