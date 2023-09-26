package models

import (
	"fmt"
	"log"

	"github.com/kkrajkumar1198/Zocket/initializers"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Product struct {
	ProductID               int            `gorm:"primary_key:auto_increment;not_null" json:"id"`
	ProductName             string         `gorm:"type:varchar(50)" json:"product_name"`
	ProductDescription      string         `gorm:"text" json:"product_desc"`
	ProductImages           pq.StringArray `gorm:"type:varchar(250)[]" json:"product_images_name"`
	ProductPrice            float64        `gorm:"size:10" json:"product_price"`
	CompressedProductImages pq.StringArray `gorm:"type:varchar(250)[]" json:"compressed_product_images_url"`
	CreatedAt               string         `json:"created_at"`
	UpdatedAt               string         `json:"updated_at"`
	DeletedAt               gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func (p Product) PrintDetails() {
	fmt.Printf("Product ID: %d\n", p.ProductID)
	fmt.Printf("Product Name: %s\n", p.ProductName)
	fmt.Printf("Product Description: %s\n", p.ProductDescription)
	fmt.Printf("Product Images: %v\n", p.ProductImages)
	fmt.Printf("Product Price: %.2f\n", p.ProductPrice)
	fmt.Printf("Compressed Product Images: %v\n", p.CompressedProductImages)
	fmt.Printf("CreatedAt: %s\n", p.CreatedAt)
	fmt.Printf("UpdatedAt: %s\n", p.UpdatedAt)
	fmt.Printf("DeletedAt: %v\n", p.DeletedAt)
}

// To update particular Product Data's CompressedFileLink
func UpdateCompressedDataLocation(id int, compressed_file_links []string) {

	log.Println("Updating the CompressedFileLink")
	var product *Product
	initializers.DB.First(&product, id)
	log.Println(product)
	initializers.DB.Model(&product).Updates(Product{
		CompressedProductImages: compressed_file_links,
		UpdatedAt:               currentTime,
	})

}
