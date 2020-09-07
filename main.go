package main

import (
	"context"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

type Post struct {
	Username    string `json:"username"`
	Jumlahlike  string `json:"jumlah like"`
	Jumlahkomen string `json:"jumlah komen"`
}

func main() {
	r := gin.Default()

	p := r.Group("/post")
	{
		p.GET("/", Get)
		p.GET("/:id", GetByID)
		p.POST("/create", Create)
		p.PUT("/:id", Update)
		p.DELETE("/:id", Delete)
	}
	r.Run(":8080")
}

func Get(c *gin.Context) {
	client, err := database.Mongodb()
	if err != nil {
		c.String(500, err.Error())
		return
	}

	result, err := client.Database("post").Collection("post").Find(context.Background(), bson.M{})
	if err != nil {
		c.String(500, err.Error())
		return
	}

	var data []map[string]interface{}
	result.All(context.Background(), &data)

	c.JSON(200, data)
}

func GetByID(c *gin.Context) {
	id := c.Param("id")

	client, err := database.Mongodb()
	if err != nil {
		c.String(500, err.Error())
		return
	}

	result, err := client.Database("post").Collection("post").Find(context.Background(), bson.M{})
	if err != nil {
		c.String(500, err.Error())
		return
	}

	var data []map[string]interface{}
	result.All(context.Background(), &data)

	c.JSON(200, data)
}

func Create(c *gin.Context) {
	var post Post
	c.BindJSON(&post)

	client, err := database.Mongodb()
	if err != nil {
		c.String(500, err.Error())
		return
	}

	_, err = client.Database("post").Collection("post").InsertOne(context.Background(), post)
	if err != nil {
		c.String(500, err.Error())
		return
	}

	c.String(200, "success")
}

func Update(c *gin.Context) {
	username := c.Param("id")

	var post Post
	c.BindJSON(&post)

	client, err := database.Mongodb()
	if err != nil {
		c.String(500, err.Error())
		return
	}

	_, err = client.Database("post").Collection("post").UpdateOne(context.Background(), bson.M{"username": username}, bson.M{"$set": post})
	if err != nil {
		c.String(500, err.Error())
		return
	}

	c.String(200, "success")
}

func Delete(c *gin.Context) {
	id := c.Param("id")

	client, err := database.Mongodb()
	if err != nil {
		c.String(500, err.Error())
		return
	}

	_, err = client.Database("post").Collection("post").DeleteOne(context.Background(), bson.M{"username": id})
	if err != nil {
		c.String(500, err.Error())
		return
	}

	c.String(200, "success")
}
