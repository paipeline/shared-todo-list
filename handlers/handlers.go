package handlers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/paipeline/todo-app/models"
	mongo "github.com/paipeline/todo-app/pkg/nosql"
	"go.mongodb.org/mongo-driver/bson"
)

const collectionName = "tasks"

// GETS

func GetTasks(c *gin.Context) {
	var tasks []models.Task
	cursor, err := mongo.FindAll(collectionName, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取任务失败"})
		return
	}
	if err = cursor.All(context.TODO(), &tasks); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "解析任务失败"})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

func GetTaskByID(c *gin.Context) {
	id := c.Param("id")
	var task models.Task
	err := mongo.FindOne(collectionName, bson.M{"id": id}).Decode(&task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取任务失败"})
		return
	}
	c.JSON(http.StatusOK, task)
}

// POSTS

func CreateTask(c *gin.Context) {
	var newTask models.Task
	if err := c.BindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newTask.ID = uuid.New()
	_, err := mongo.InsertOne(collectionName, newTask)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建任务失败"})
		return
	}
	c.JSON(http.StatusCreated, newTask)
}

func UpdateTaskByID(c *gin.Context) {
	id := c.Param("id")
	var updatedTask models.Task
	if err := c.BindJSON(&updatedTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	update := bson.M{"$set": bson.M{"completed": updatedTask.Completed}}
	result, err := mongo.UpdateOne(collectionName, bson.M{"id": id}, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "更新任务失败"})
		return
	}
	if result.MatchedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "任务未找到"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "任务已更新"})
}

func DeleteTaskByID(c *gin.Context) {
	id := c.Param("id")
	result, err := mongo.DeleteOne(collectionName, bson.M{"id": id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "删除任务失败"})
		return
	}
	if result.DeletedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "任务未找到"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "任务已删除"})
}
