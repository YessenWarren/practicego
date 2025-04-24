package test

import (
	"sneakers/database"
	"log"
)

func setupDatabase() {
	database.InitDB() // Удалили := err и проверку на ошибку

	// Очистка таблиц
	if err := database.DB.Exec("DELETE FROM sneakers").Error; err != nil {
		log.Fatalf("Failed to clean sneakers table: %v", err)
	}
	if err := database.DB.Exec("DELETE FROM brands").Error; err != nil {
		log.Fatalf("Failed to clean brands table: %v", err)
	}
	if err := database.DB.Exec("DELETE FROM categories").Error; err != nil {
		log.Fatalf("Failed to clean categories table: %v", err)
	}
	if err := database.DB.Exec("DELETE FROM users").Error; err != nil {
		log.Fatalf("Failed to clean users table: %v", err)
	}
}
