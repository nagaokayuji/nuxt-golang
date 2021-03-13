package controllers

import (
	"fmt"
	"net/http"

	"app/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// 全件取得
func GetAllTodos(c *gin.Context) {
	var todos []models.Todo
	models.DB.Find(&todos)
	c.JSON(http.StatusOK, todos)
}

// UUIDから1件取得
func GetTodo(c *gin.Context) {
	uuid := c.Params.ByName("uuid")
	var todo models.Todo
	if err := models.DB.Where("uuid = ?", uuid).First(&todo).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func CreateTodo(c *gin.Context) {
	// TODO not yet implemented
	fmt.Println(uuid.New().String())
}
