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

func ReadAllPosts(c *gin.Context) {
	var DATABASE = database.ConnectDB()
	var postCollection = collection.GetCollection(DATABASE, "Posts")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	postEmail := c.Param("email")
	defer cancel()

	info, err := postCollection.Find(ctx, bson.D{primitive.E{Key: "email", Value: postEmail}})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "No se encontro nada con ese titulo"})
		return
	}

	var results []model.Posts
	if err = info.All(ctx, &results); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Algo salio mal al intentar pasar los info a la variable result"})
		return
	}

	check := len(results)
	if check == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "No hay coincidencias"})
	} else {
		c.JSON(http.StatusCreated, results)
	}
}
