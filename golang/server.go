package main

import (
	"app/controllers"

	"app/models"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	// 起動
	server()
}

func server() {
	db := models.SetupDB()
	db.AutoMigrate(&models.Todo{})
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world from api",
		})
	})

	// 全件
	r.GET("/todos", controllers.GetAllTodos)

	r.Run(":8080")
}
