package mongo

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client *mongo.Client
	db     *mongo.Database
)

func ConnectDB() error {
	log.Println("正在连接到MongoDB...")
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	var err error
	client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Printf("连接MongoDB失败: %v", err)
		return err
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Printf("Ping MongoDB失败: %v", err)
		return err
	}
	log.Println("成功连接到MongoDB!")
	return nil
}

func CloseDB() {
	log.Println("正在断开MongoDB连接...")
	if client != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		err := client.Disconnect(ctx)
		if err != nil {
			log.Printf("断开MongoDB连接时发生错误: %v", err)
		} else {
			log.Println("成功断开MongoDB连接!")
		}
	}
}

// 初始化数据库
func InitDB(dbName string) error {
	log.Printf("正在初始化数据库: %s", dbName)
	if client == nil {
		return fmt.Errorf("MongoDB客户端未初始化")
	}
	db = client.Database(dbName)
	log.Printf("成功初始化数据库: %s", dbName)
	return nil
}

// 插入文档
func InsertOne(collectionName string, document interface{}) (*mongo.InsertOneResult, error) {
	log.Printf("正在插入文档到集合: %s", collectionName)
	collection := db.Collection(collectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result, err := collection.InsertOne(ctx, document)
	if err != nil {
		log.Printf("插入文档失败: %v", err)
	} else {
		log.Printf("成功插入文档，ID: %v", result.InsertedID)
	}
	return result, err
}

// 查找文档
func FindOne(collectionName string, filter bson.M) *mongo.SingleResult {
	log.Printf("正在查找文档，集合: %s, 过滤条件: %v", collectionName, filter)
	collection := db.Collection(collectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return collection.FindOne(ctx, filter)
}

// 更新文档
func UpdateOne(collectionName string, filter bson.M, update bson.M) (*mongo.UpdateResult, error) {
	log.Printf("正在更新文档，集合: %s, 过滤条件: %v, 更新内容: %v", collectionName, filter, update)
	collection := db.Collection(collectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Printf("更新文档失败: %v", err)
	} else {
		log.Printf("成功更新文档，匹配数: %d, 修改数: %d", result.MatchedCount, result.ModifiedCount)
	}
	return result, err
}

// 删除文档
func DeleteOne(collectionName string, filter bson.M) (*mongo.DeleteResult, error) {
	log.Printf("正在删除文档，集合: %s, 过滤条件: %v", collectionName, filter)
	collection := db.Collection(collectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		log.Printf("删除文档失败: %v", err)
	} else {
		log.Printf("成功删除文档，删除数: %d", result.DeletedCount)
	}
	return result, err
}

func FindAll(collectionName string, filter bson.M) (*mongo.Cursor, error) {
	log.Printf("正在查找所有文档，集合: %s, 过滤条件: %v", collectionName, filter)
	collection := db.Collection(collectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		log.Printf("查找所有文档失败: %v", err)
	} else {
		log.Println("成功查找所有文档")
	}
	return cursor, err
}
