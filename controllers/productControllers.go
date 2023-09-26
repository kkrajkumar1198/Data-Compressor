package controllers

import (
	"log"
	"time"

	"github.com/kkrajkumar1198/Zocket/initializers"
	"github.com/kkrajkumar1198/Zocket/models"
)

var (
	currentTime = time.Now().Format("2006-1-2 15:04:05")
)

func CreateProduct(prod models.Product) {
	prod.CreatedAt = currentTime
	prod.UpdatedAt = currentTime
	result := initializers.DB.Create(&prod)
	if result.Error != nil {
		log.Printf("Create Error: %s", result.Error)
	}

}
