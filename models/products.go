package models

import "gorm.io/gorm"

type Product struct {
	ProductID               int
	ProductName             string
	ProductDescription      int64
	ProductImages           []string
	ProductPrice            float64
	CompressedProductImages []string
	CreateAt                string
	UpdatedAt               string
	DeletedAt               gorm.DeletedAt `gorm:"index"`
}
