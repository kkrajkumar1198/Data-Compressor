package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kkrajkumar1198/Zocket/controllers"
	"github.com/kkrajkumar1198/Zocket/initializers"
	"github.com/kkrajkumar1198/Zocket/messagequeue"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.GetGCSClient()
	initializers.StartKafka()
	messagequeue.Start()

}
func main() {
	r := gin.Default()

	// Crud Operations on User
	r.POST("/createUser", controllers.CreateUser)
	r.GET("/getUser/:name", controllers.GetUserData)
	r.GET("/getUsers", controllers.GetUsers)
	r.PUT("/updateUser/:id", controllers.UpdateUser)
	r.DELETE("/deleteUser/:id", controllers.DeleteUser)

	// This POST operation will call the producer method in Kafka
	// Producer - saves the product data to DB and sends ProductID to Consumer
	// Consumer - Consumes fetch Product data using ProductID and download data from GCS and compress and store it in local

	r.POST("/producer", messagequeue.NewProduct)
	r.Run()
}
