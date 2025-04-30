package handlers

import (
	"net/http"
	"sneakers/database"
	"sneakers/models"
	"github.com/gin-gonic/gin"
)

// Получение списка отзывов с пагинацией
func GetReviews(c *gin.Context) {
	var reviews []models.Review
	if err := database.DB.Find(&reviews).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, reviews)
}

// Создание нового отзыва
func CreateReview(c *gin.Context) {
	var review models.Review
	if err := c.ShouldBindJSON(&review); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректные данные"})
		return
	}

	if err := database.DB.Create(&review).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при создании отзыва"})
		return
	}

	c.JSON(http.StatusCreated, review)
}

// Получение отзыва по ID
func GetReviewByID(c *gin.Context) {
	var review models.Review
	if err := database.DB.First(&review, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Отзыв не найден"})
		return
	}
	c.JSON(http.StatusOK, review)
}

// Обновление отзыва
func UpdateReview(c *gin.Context) {
	var review models.Review
	if err := database.DB.First(&review, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Отзыв не найден"})
		return
	}

	if err := c.ShouldBindJSON(&review); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректные данные"})
		return
	}

	if err := database.DB.Save(&review).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при обновлении отзыва"})
		return
	}

	c.JSON(http.StatusOK, review)
}

// Удаление отзыва
func DeleteReview(c *gin.Context) {
	result := database.DB.Delete(&models.Review{}, c.Param("id"))
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при удалении отзыва"})
		return
	}
	if result.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Отзыв не найден"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Отзыв удалён"})
}
