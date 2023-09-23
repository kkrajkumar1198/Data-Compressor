package main

import (
	initializers "github.com/kkrajkumar1198/Zocket/Initializers"
	"github.com/kkrajkumar1198/Zocket/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Users{})
}
