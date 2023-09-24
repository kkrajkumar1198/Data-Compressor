package main

import (
	"github.com/kkrajkumar1198/Zocket/initializers"
	"github.com/kkrajkumar1198/Zocket/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.User{})
	initializers.DB.AutoMigrate(&models.Product{})
}
