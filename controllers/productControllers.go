package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/kkrajkumar1198/Zocket/initializers"
	"github.com/kkrajkumar1198/Zocket/models"
)

func CreateProduct(c *gin.Context) {

	var data models.Product
	// Get Data struct
	c.Bind(&data)

	product := models.Product{ProductName: data.ProductName, ProductDescription: data.ProductDescription, ProductImages: data.ProductImages, ProductPrice: data.ProductPrice}

	result := initializers.DB.Create(&product)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"product": product,
	})
}

// func UpdateUser(c *gin.Context) {

// 	id := c.Param("id")

// 	var user models.Users
// 	initializers.DB.First(&user, id)
// 	// Get Data struct
// 	c.Bind(&user)

// 	initializers.DB.Model(&user).Updates(models.Users{
// 		Name:       user.Name,
// 		Mobile_num: user.Mobile_num,
// 		Latitude:   user.Latitude,
// 		Longitude:  user.Longitude,
// 	})

// 	c.JSON(200, gin.H{
// 		"user": user,
// 	})
// }

// func GetUsers(c *gin.Context) {
// 	var users []models.Users
// 	initializers.DB.Find(&users)

// 	c.JSON(200, gin.H{
// 		"user": users,
// 	})
// }

// func GetUser(c *gin.Context) {

// 	// Get Id
// 	id := c.Param("id")

// 	var user models.Users
// 	initializers.DB.First(&user, id)

// 	c.JSON(200, gin.H{
// 		"user": user,
// 	})
// }

// func DeleteUser(c *gin.Context) {

// 	id := c.Param("id")
// 	initializers.DB.Delete(&models.Users{}, id)
// 	c.Status(200)
// }
