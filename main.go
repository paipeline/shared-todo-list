package main

import (
	"github.com/gin-gonic/gin"
	"github.com/paipeline/todo-app/handlers"
	mongo "github.com/paipeline/todo-app/pkg/nosql"
)

func main() {
	// 连接到 MongoDB
	mongo.ConnectDB()
	defer mongo.CloseDB()

	// 初始化数据库
	mongo.InitDB("todo_app")

	router := gin.Default()
	router.GET("/tasks", handlers.GetTasks)
	router.POST("/tasks", handlers.CreateTask)
	router.GET("/tasks/:id", handlers.GetTaskByID) // 注意这里修正了URL参数格式
	router.PUT("/tasks/:id", handlers.UpdateTaskByID)
	router.DELETE("/tasks/:id", handlers.DeleteTaskByID)

	router.Run("localhost:8080")
}
