package controllers

import (
	"fmt"
	"net/http"

	"app/models"
	"time"

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
		return
	}

	c.JSON(http.StatusOK, todo)
}

// 1件登録
func CreateTodo(c *gin.Context) {
	var input models.CreateTodoInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println(err.Error())
		return
	}

	todo := models.Todo{UUID: uuid.New().String(), Title: input.Title, State: false, Deadline: input.Deadline, CreatedAt: time.Now(), UpdatedAt: time.Now()}
	models.DB.Create(&todo)
	c.JSON(http.StatusOK, todo)
}
