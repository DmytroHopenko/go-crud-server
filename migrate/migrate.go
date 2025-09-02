package main

import (
	initializers "github.com/DmytroHopenko/go-crud-server/initializerz"
	"github.com/DmytroHopenko/go-crud-server/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
