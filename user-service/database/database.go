package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	// Fetch environment variables for MySQL credentials
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = "mysql-1"
	}

	// Construct DSN (Data Source Name) for MySQL connection
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbName)

	// Open a connection to the database using GORM
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database, got error: %v", err)
	}

	// Store the DB instance for later use
	DB = db
	fmt.Println("Successfully connected to the database")
}
