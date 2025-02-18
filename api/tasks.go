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

func ListTask(c *gin.Context) {
	var listTask service.ListTaskService
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&listTask); err != nil {
		logging.Error(err)
		c.JSON(400, err)
	} else {
		res := listTask.List(claim.Id)
		c.JSON(200, res)
	}
}

func UpdateTask(c *gin.Context) {
	var updateTask service.UpdateTaskService
	if err := c.ShouldBind(&updateTask); err != nil {
		logging.Error(err)
		c.JSON(400, err)
	} else {
		res := updateTask.Update(c.Param("id"))
		c.JSON(200, res)
	}
}

func SearchTask(c *gin.Context) {
	var searchTask service.SearchTaskService
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&searchTask); err != nil {
		logging.Error(err)
		c.JSON(400, err)
	} else {
		res := searchTask.Search(claim.Id)
		c.JSON(200, res)
	}
}

func DeleteTask(c *gin.Context) {
	var deleteTask service.DeleteTaskService
	if err := c.ShouldBind(&deleteTask); err != nil {
		logging.Error(err)
		c.JSON(400, err)
	} else {
		res := deleteTask.Delete(c.Param("id"))
		c.JSON(200, res)
	}
}
