package models

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	Name       string
	Mobile_num int16
	Latitude   float64
	Longitude  float64
}
