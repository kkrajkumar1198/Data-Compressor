# Product Data Processor
 ### The **Product Data Processor** is a robust and efficient solution designed to manage product-related data in a distributed environment. It seamlessly handles the storage, processing, and retrieval of product information, making it a vital component for businesses dealing with a diverse range of products.
---

## Tech Stack

The following technologies and libraries have been used in this project:

- **Programming Language**: Go (Golang)
- **Web Framework**: Gin (for API handling)
- **Database**: PostgreSQL -- ElephantSQL (PostgreSQL as API)
- **Cloud Storage**: Google Cloud Storage (GCS)
- **Message Queue**: Kafka
- **Object-Relational Mapping (ORM)**: GORM
- **Testing Framework**: Testing using `testing` package
---

## Setup:
- Download Go
- Do go mod tidy, it will install all the required packages
- CompileDaemon -command="./Zocket" - Start CompileDaemon to run the project

## For Integration testing:
 - go test -timeout 30s -run ^TestDBConnectionIntegration$ --- *DB Connection Test*
 - go test -timeout 30s -run ^TestGCSClientIntegration$ --- *GCS Connection Test*
 - go test -timeout 30s -run ^TestKafkaIntegration$ --- *Kafka Connection Test*

## For Benchmark testing:
 - go test -benchmem -run=^$ -bench ^BenchmarkDownloadAndCompressImages$  --- *Download and Compression Test*

## Database Schema

### Users Table


| Column Name | Data Type       | Description                          |
|-------------|-----------------|--------------------------------------|
| ID          | int (Primary Key)| Unique user identifier              |
| Name        | varchar(50)     | User's name                          |
| MobileNumber      | int             | Contact number of the user           |
| Latitude    | float64         | Latitude of the user's location      |
| Longitude   | float64         | Longitude of the user's location     |
| DateJoined  | timestamp       | Timestamp of user creation           |
| UpdatedAt  | timestamp       | Timestamp of last update             |
| DeletedAt  | timestamp       | Timestamp of user deletion data             |

### Products Table


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

| Column Name               | Data Type       | Description                             |
|---------------------------|-----------------|-----------------------------------------|
| ProductID                | int (Primary Key)| Unique product identifier              |
| ProductName              | varchar(50)     | Name of the product                     |
| ProductDescription       | text            | Description of the product              |
| ProductImages           | varchar(250)[]   | Array of image URLs associated with the product |
| ProductPrice             | float64  | Price of the product                   |
| CompressedProductImages | varchar(250)[]  | Array of compressed image locations    |
| CreatedAt                | timestamp       | Timestamp of product creation          |
| DeletedAt  | timestamp       | Timestamp of product deletion data             |

---

## API Documentation
 - Please refer Data Processor.postman_collection.json - Postman Collection File 
 - Import the file in Postman to do API Calls


## Sensitive Information:
 - GCS Key and Postgres API key is not provided with this repo.