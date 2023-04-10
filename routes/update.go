package routes

import (
	"context"
	"golangGinMongo/collection"
	"golangGinMongo/database"
	"golangGinMongo/model"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdatePost(c *gin.Context) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var DATABASE = database.ConnectDB()
	var postCollection = collection.GetCollection(DATABASE, "Posts")

	postId := c.Param("postId")
	var post model.Posts
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(postId)

	if err := c.BindJSON(&post); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	postUpdated := new(model.Posts)

	edited := bson.M{"title": post.Title, "article": post.Article, "email": post.Email, "done": post.Done}
	result, err := postCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": edited})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	if result.MatchedCount < 1 {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "No se actualizo nada"})
		return
	}

	ok := postCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&postUpdated)
	if ok != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Post no encontrado"})
		return
	}

	c.JSON(http.StatusCreated, postUpdated)

}
