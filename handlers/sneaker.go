package handlers

import (
	"net/http"
	"sneakers/database"
	"sneakers/models"

	"github.com/gin-gonic/gin"
)

func GetSneakers(c *gin.Context) {
	var sneakers []models.Sneaker

	title := c.Query("title")

	query := database.DB

	// Фильтрация по полю Name
	if title != "" {
		query = query.Where("name ILIKE ?", "%"+title+"%")
	}

	if err := query.Find(&sneakers).Error; err != nil {
		// Возвращаем точную ошибку для отладки
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, sneakers)
}

func GetSneakerByID(c *gin.Context) {
	var sneaker models.Sneaker
	if err := database.DB.Preload("Brand").Preload("Category").First(&sneaker, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Кроссовка не найдена"})
		return
	}
	c.JSON(http.StatusOK, sneaker)
}

func CreateSneaker(c *gin.Context) {
	var sneaker models.Sneaker
	if err := c.ShouldBindJSON(&sneaker); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := database.DB.Create(&sneaker).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при создании записи"})
		return
	}
	c.JSON(http.StatusCreated, sneaker)
}

func UpdateSneaker(c *gin.Context) {
	var sneaker models.Sneaker
	if err := database.DB.First(&sneaker, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Кроссовка не найдена"})
		return
	}
	if err := c.ShouldBindJSON(&sneaker); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := database.DB.Save(&sneaker).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при обновлении записи"})
		return
	}
	c.JSON(http.StatusOK, sneaker)
}

func DeleteSneaker(c *gin.Context) {
	if err := database.DB.Delete(&models.Sneaker{}, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Не удалось удалить кроссовку"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Кроссовка удалена"})
}
