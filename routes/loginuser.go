package routes

import (
	"context"
	"golangGinMongo/collection"
	"golangGinMongo/database"
	"golangGinMongo/model"
	"golangGinMongo/tokengenerator"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

type CookieSearched struct {
	Name  string
	Value string
}

func LoginUser(c *gin.Context) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	//logeo de usuario por token y cookies

	cookieClient, error := c.Request.Cookie("CheckAuth")
	if error == nil && cookieClient.Value != "" {
		defer cancel()
		var cookie CookieSearched
		c.BindJSON(&cookie)
		if cookie.Value != cookieClient.Value {
			c.JSON(http.StatusBadRequest, gin.H{"message": "No coinciden los datos de las cookies"})
			return
		}
		token, notoken := c.Request.Cookie("Auth")
		if notoken != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "La cookie Auth no se encontro con valores o no se pudo leer"})
			return
		}
		data, err := tokengenerator.VerifyToken(token.Value)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Claims no encontrados "})
		} else {
			c.JSON(http.StatusAccepted, data)
			return
		}
	}

	//logeo normal
	var DATABASE = database.ConnectDB()
	var userCollection = collection.GetUserCollection(DATABASE, "Users")
	var user model.User
	defer cancel()

	var item bson.M
	if err := c.BindJSON(&item); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	err := userCollection.FindOne(ctx, bson.M{"email": item["email"], "password": item["password"]}).Decode(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Mail o contraseña incorrecta"})
		return
	}

	token, err := tokengenerator.GenerateToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Token no generado, algo fallo"})
		return
	}

	c.SetSameSite(http.SameSiteNoneMode)
	c.SetCookie("Auth", token, 24000, "/", "https://task-app-amvt.onrender.com", true, true)
	c.SetCookie("CheckAuth", "SiAutentico", 24000, "/", "https://task-app-amvt.onrender.com", true, false)
	c.JSON(http.StatusCreated, user)
}
