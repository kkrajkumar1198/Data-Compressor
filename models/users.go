package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

var (
	currentTime = time.Now().Format("2006-1-2 15:04:05")
)

type User struct {
	ID           int     `gorm:"primary_key:auto_increment;not_null" json:"id"`
	Name         string  `gorm:"type:varchar(50)" json:"name"`
	MobileNumber int     `gorm:"size:100" json:"mobile_number"`
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
	DateJoined   string
	UpdatedAt    string
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

// PrintUserData is a method for the User struct to print all its data.
func (u User) PrintUserData() {
	fmt.Printf("ID: %d\n", u.ID)
	fmt.Printf("Name: %s\n", u.Name)
	fmt.Printf("Mobile Number: %d\n", u.MobileNumber)
	fmt.Printf("Latitude: %f\n", u.Latitude)
	fmt.Printf("Longitude: %f\n", u.Longitude)
	fmt.Printf("Date Joined: %s\n", u.DateJoined)
	fmt.Printf("Updated At: %s\n", u.UpdatedAt)
	fmt.Printf("Deleted At: %s\n", u.DeletedAt.Time.String())
}

// To fetch All Users Data
func GetAllUser(db *gorm.DB, user *[]User) (err error) {
	err = db.Find(user).Error
	if err != nil {
		return err
	}
	return nil
}

// To create a new user
func CreateUser(db *gorm.DB, data *User) (err error) {

	data.DateJoined = currentTime
	data.UpdatedAt = currentTime

	result := db.Create(&data)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// To fetch particular User Data by name
func GetByUsername(db *gorm.DB, data *User, username string) (err error) {

	err = db.Where("name = ?", username).First(data).Error

	if err != nil {
		return err
	}
	return nil
}

// To update particular User Data by id
func UpdateUserData(db *gorm.DB, data *User, id string) {

	var user *User
	db.First(&user, id)

	db.Model(&user).Updates(User{
		Name:         data.Name,
		MobileNumber: data.MobileNumber,
		Latitude:     data.Latitude,
		Longitude:    data.Longitude,
		UpdatedAt:    currentTime,
	})

}

/*
EMBEDDED STRUCT:
We can use gorm.Model which will auto import default attributes
type User struct {
	gorm.Model
	Name string
  }
  // equals
  type User struct {
	ID        uint           `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Name string
}
*/
