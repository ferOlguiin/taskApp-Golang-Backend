package main

import (
	"golangGinMongo/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	//cors
	config := cors.DefaultConfig()

	//cors config
	config.AllowOrigins = []string{"http://localhost:5173"}

	//estas credenciales van a ir a parar al frontend donde van a alojar datos en las cookies
	config.AllowCredentials = true

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
	router.Run("localhost:3000")
}
