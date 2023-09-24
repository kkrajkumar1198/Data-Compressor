package initializers

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	//postgres://jgkaapdo:6V_h7nwZux95yanKNQFBxoVKpB5ABn99@bubble.db.elephantsql.com/jgkaapdo
	dsn := os.Getenv("DB_URL")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database.")
	}
}
