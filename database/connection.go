package database

import (
	"log"
	"os"
	"todo-app/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
    if err := godotenv.Load(); err != nil {
        log.Fatal("Error loading .env file")
    }

    dsn := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASS") + "@tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/" + os.Getenv("DB_NAME") + "?charset=utf8mb4&parseTime=True&loc=Local"
    
    var err error
    
    DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database", err)
    }

    // Migrate the schema
    DB.AutoMigrate(&models.Todo{})

    if err := godotenv.Load(); err != nil {
        log.Fatal("Error loading .env file")
    }
}
