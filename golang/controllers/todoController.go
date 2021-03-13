package controllers

import (
	"net/http"

	"app/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func GetAllTodos(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var todos []models.Todo
	db.Find(&todos)
	c.JSON(http.StatusOK, gin.H{"data": todos})
}
