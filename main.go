package main

import (
	"golangGinMongo/routes"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	PORT := os.Getenv("PORT")
	if PORT == "" {
		log.Fatal("port no puede ser un string vacio!")
	}

	//cors
	config := cors.DefaultConfig()

	//cors config
	config.AllowOrigins = []string{"https://task-app-amvt.onrender.com"}

	//estas credenciales van a ir a parar al frontend donde van a alojar datos en las cookies
	config.AllowCredentials = true
	config.AllowHeaders = []string{"Origin", "Content-Type", "withCredentials"}
	config.AllowMethods = []string{"GET", "POST", "DELETE", "PUT", "PATCH", "OPTIONS"}

	//use cors
	router.Use(cors.New(config))

	//routes
	router.POST("/post", routes.CreatePost)
	router.GET("/post/:postId", routes.ReadOnePost)
	router.DELETE("/post/:postId", routes.DeletePost)
	router.PUT("/post/:postId", routes.UpdatePost)
	router.GET("/allpost/:email", routes.ReadAllPosts)
	router.POST("/user", routes.LoginUser)
	router.POST("/createuser", routes.CreateUser)
	router.DELETE("/user/:userid", routes.DeleteUser)

	//server
	router.Run("0.0.0.0:" + PORT)
}
