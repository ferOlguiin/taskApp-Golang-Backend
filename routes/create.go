package routes

import (
	"context"
	"golangGinMongo/collection"
	"golangGinMongo/database"
	"golangGinMongo/model"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func CreatePost(c *gin.Context) {

	var DATABASE = database.ConnectDB()
	var postCollection = collection.GetCollection(DATABASE, "Posts")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	post := new(model.Posts)

	defer cancel()

	if err := c.BindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		log.Fatal(err)
		return
	}
	postPayload := model.Posts{
		Title:   post.Title,
		Article: post.Article,
		Email:   post.Email,
		Done:    post.Done,
	}

	result, err := postCollection.InsertOne(ctx, postPayload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	postFound := new(model.Posts)
	ok := postCollection.FindOne(ctx, bson.M{"_id": result.InsertedID}).Decode(&postFound)
	if ok != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "El post con ese id no existe"})
		return
	}

	c.JSON(http.StatusCreated, postFound)

}
