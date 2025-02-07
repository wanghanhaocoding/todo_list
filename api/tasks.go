package api

import (
	"github.com/gin-gonic/gin"
	logging "github.com/sirupsen/logrus"
	"todo_list/pkg/utils"
	"todo_list/service"
)

func CreateTask(c *gin.Context) {
	var createTask service.CreateTaskService
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&createTask); err != nil {
		logging.Error(err)
		c.JSON(400, err)
	} else {
		res := createTask.Create(claim.Id)
		c.JSON(200, res)
	}
}

func ShowTask(c *gin.Context) {
	var showTask service.ShowTaskService
	if err := c.ShouldBind(&showTask); err != nil {
		logging.Error(err)
		c.JSON(400, err)
	} else {
		res := showTask.Show(c.Param("id"))
		c.JSON(200, res)
	}
}
