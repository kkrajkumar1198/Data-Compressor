package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kkrajkumar1198/Zocket/cloudbucket"
	"github.com/kkrajkumar1198/Zocket/controllers"
	"github.com/kkrajkumar1198/Zocket/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}
func main() {
	r := gin.Default()
	r.POST("/createUser", controllers.CreateUser)
	r.GET("/getUser/:name", controllers.GetUserData)
	r.GET("/getUsers", controllers.GetUsers)
	r.PUT("/updateUser/:id", controllers.UpdateUser)
	// r.DELETE("/deleteUser/:id", controllers.DeleteUser)
	r.GET("/downloadImage", cloudbucket.DownloadAndSaveImageToLocal)
	r.Run()
}
