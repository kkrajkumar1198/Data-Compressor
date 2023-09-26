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
 - go test -benchmem -run=^$ -bench ^BenchmarkDownloadAndCompressImages$  --- *DB Connection Test*

## Database Schema

### Users Table

| Column Name | Data Type       | Description                          |
|-------------|-----------------|--------------------------------------|
| id          | int (Primary Key)| Unique user identifier              |
| name        | varchar(50)     | User's name                          |
| mobile      | int             | Contact number of the user           |
| latitude    | float64         | Latitude of the user's location      |
| longitude   | float64         | Longitude of the user's location     |
| created_at  | timestamp       | Timestamp of user creation           |
| updated_at  | timestamp       | Timestamp of last update             |

### Products Table

| Column Name               | Data Type       | Description                             |
|---------------------------|-----------------|-----------------------------------------|
| product_id                | int (Primary Key)| Unique product identifier              |
| product_name              | varchar(50)     | Name of the product                     |
| product_description       | text            | Description of the product              |
| product_images           | varchar(250)[]   | Array of image URLs associated with the product |
| product_price             | numeric(10, 2)  | Price of the product                   |
| compressed_product_images | varchar(250)[]  | Array of compressed image locations    |
| created_at                | timestamp       | Timestamp of product creation          |
| updated_at                | timestamp       | Timestamp of last update               |

---

## API Documentation
 - Please refer Data Processor.postman_collection.json - Postman Collection File 
 - Import the file in Postman to do API Calls


## Sensitive Information:
 - GCS Key and Postgres API key is not provided with this repo.