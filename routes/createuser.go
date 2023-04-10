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

func CreateUser(c *gin.Context) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	var DATABASE = database.ConnectDB()
	var userCollection = collection.GetUserCollection(DATABASE, "Users")

	var user model.User

	defer cancel()

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		log.Fatal(err)
		return
	}

	newUser := model.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}

	//busco que el email no este registrado por otro usuario para crear uno nuevo
	var item bson.M
	err := userCollection.FindOne(ctx, bson.M{"email": newUser.Email}).Decode(&item)
	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "El email que intenta registrar ya existe"})
		return
	}

	//si el email no esta registrado se procede a crear uno nuevo
	result, err := userCollection.InsertOne(ctx, newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Usuario creado correctamente",
		"data":    result,
		"item":    newUser,
	})

}
