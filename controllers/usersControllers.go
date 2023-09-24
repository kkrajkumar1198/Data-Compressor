package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kkrajkumar1198/Zocket/initializers"
	"github.com/kkrajkumar1198/Zocket/models"
)

func CreateUser(c *gin.Context) {

	var data models.User
	// Get Data struct
	c.Bind(&data)

	err := models.CreateUser(initializers.DB, &data)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
	}
	c.JSON(http.StatusOK, "User Added Successfully")

}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	// Get Data struct
	c.Bind(&user)
	models.UpdateUserData(initializers.DB, &user, id)
	c.JSON(200, gin.H{
		"user": user,
	})
}

func GetUsers(c *gin.Context) {
	var users []models.User
	err := models.GetAllUser(initializers.DB, &users)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
	}
	c.JSON(http.StatusOK, users)
}

func GetUserData(c *gin.Context) {

	// Get Id
	name := c.Param("name")

	var user models.User
	err := models.GetByUsername(initializers.DB, &user, name)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
	}
	c.JSON(http.StatusOK, user)
}

// func DeleteUser(c *gin.Context) {

// 	id := c.Param("id")
// 	initializers.DB.Delete(&models.User{}, id)
// 	c.Status(200)
// }
