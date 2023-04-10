package routes

import (
	"context"
	"golangGinMongo/collection"
	"golangGinMongo/database"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteUser(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	var DATABASE = database.ConnectDB()
	id := c.Param("userid")
	var collection = collection.GetUserCollection(DATABASE, "Users")
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(id)
	result, err := collection.DeleteOne(ctx, bson.M{"_id": objId})
	res := map[string]interface{}{"data": result}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "No hay ningún usuario con ese id"})
		return
	}

	if result.DeletedCount < 1 {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "No se borro nada, el contador de datos borrados esta en 0"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Usuario borrado correctamente", "data": res})

}
