package initializers

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Connect to DB
func ConnectDB() {
	var err error
	dsn := os.Getenv("DB_DSN")
	// store the DB connection in DB variable
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to DB")
	} else {
		log.Println("Connected to DB successfully")
	}

}

// Get DB connection
func GetDB() *gorm.DB {
	// return the DB connection
	return DB
}
