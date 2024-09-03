package main

import (
	"github.com/gin-gonic/gin"
	"github.com/paipeline/todo-app/handlers" // 使用您的项目路径
)

func main() {
	router := gin.Default()
	router.GET("/tasks", handlers.GetTasks)
	router.POST("/tasks", handlers.CreateTask)
	router.GET("/tasks/:id", handlers.GetTaskByID) // 注意这里修正了URL参数格式
	router.PUT("/tasks/:id", handlers.UpdateTaskByID)
	router.DELETE("/tasks/:id", handlers.DeleteTaskByID)

	router.Run("localhost:8080")
}
