package handlers

import (
	"net/http"

	"github.com/paipeline/todo-app/models" // 导入models包

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// GETS

func GetTasks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, models.Tasks)
}

func GetTaskByID(c *gin.Context) {
	id := c.Param("id")
	for _, task := range models.Tasks {
		if task.ID.String() == id {
			c.IndentedJSON(http.StatusOK, task)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Task not found"})
}

// POSTS

func CreateTask(c *gin.Context) {
	var newTask models.Task

	if err := c.BindJSON(&newTask); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newTask.ID = uuid.New()
	models.Tasks = append(models.Tasks, newTask)
	c.IndentedJSON(http.StatusCreated, newTask)
}

func UpdateTaskByID(c *gin.Context) {
	id := c.Param("id")
	for index, task := range models.Tasks {
		if task.ID.String() == id {
			models.Tasks[index].Completed = true
			c.IndentedJSON(http.StatusOK, models.Tasks[index])
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Task not found"})
}

func DeleteTaskByID(c *gin.Context) {
	id := c.Param("id")
	for index, task := range models.Tasks {
		if task.ID.String() == id {
			models.Tasks = append(models.Tasks[:index], models.Tasks[index+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "Task deleted"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Task not found"})
}
