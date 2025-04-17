package database

import (
	"fmt"
	"log"
    "sneakers/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "host=localhost user=postgres password=20328 dbname=sneakers port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	fmt.Println("Database connected successfully!")

	err = DB.AutoMigrate(&models.Sneaker{}, &models.Brand{}, &models.Category{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
}
