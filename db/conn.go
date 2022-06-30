package db

import (
	"go_jwt_samplev2/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
func Connect() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	a := os.Getenv("dbConn")

	conn, err := gorm.Open(mysql.Open(a), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	DB = conn
	conn.AutoMigrate(&models.User{}) //Auto created
}
