package db

import (
	"go_jwt_samplev2/models"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

	dbConnInfo := os.Getenv("DB_CONN")
	conn, err := gorm.Open(mysql.Open(dbConnInfo), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	DB = conn
	conn.AutoMigrate(&models.User{}) //Auto created
}
