package main

import (
	"github.com/gin-gonic/gin"
	"sneakers/database"
	"sneakers/models"
	"sneakers/routes"
)

func main() {
	database.InitDB()

	database.DB.AutoMigrate(&models.User{}, &models.Sneaker{}, &models.Brand{}, &models.Category{})

	r := gin.Default()
	routes.SetupRoutes(r)

	// 👇 Указываем порт 9090 явно
	r.Run(":9999")
}
