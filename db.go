package main

import(
	"log"
	"os"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *gorm.DB

func InitDB(){
	// Load env variables from .env file
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Read the database URL from env variables
	databaseURL := os.Getenv("DATABASE_URL")

	var err error

	DB, err = gorm.Open("postgres", databaseURL)

	if err != nil {
		log.Fatalf("Failed to connect database: %w", err)
	}

	// Create table if not exists
	DB.AutoMigrate(&URL{}).Error

	if err != nil {
		log.Fatalf("Error during migration %v", err)
	}

	log.Println("Successfully connected to the database and migrated tables")


}
