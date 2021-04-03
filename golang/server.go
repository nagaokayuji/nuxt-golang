package main

import (
	"golang/controllers"

	"golang/models"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	// 起動
	server()
}

func server() {
	models.SetupDB()
	models.DB.AutoMigrate(&models.Todo{})

	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Set("db", models.DB)
	})

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world from api",
		})
	})

	// 全件
	r.GET("/todos", controllers.GetAllTodos)
	// 1件
	r.GET("/todos/:uuid", controllers.GetTodo)
	// 1件作成
	r.POST("/todos", controllers.CreateTodo)
	// toggle
	r.PATCH("/todos/:uuid", controllers.ToggleStatus)

	r.Run(":8080")
}
