package main

import (
	"github.com/DmytroHopenko/go-crud-server/controllers"
	initializers "github.com/DmytroHopenko/go-crud-server/initializerz"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {
	router := gin.Default()

	router.POST("/posts", controllers.PostsCreate)
	router.PUT("/posts/:id", controllers.PostUpdate)
	router.GET("/posts", controllers.PostsIndex)
	router.GET("/posts/:id", controllers.PostShow)
	router.DELETE("/posts/:id", controllers.PostDelete)

	router.Run() // listen and serve on 0.0.0.0:8080
}
