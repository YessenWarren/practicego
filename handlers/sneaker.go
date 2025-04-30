package handlers

import (
	"net/http"
	"sneakers/database"
	"sneakers/models"
    "strconv"
	"github.com/gin-gonic/gin"
)

func GetSneakers(c *gin.Context) {
	var sneakers []models.Sneaker

	// Пагинация
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		page = 1
	}
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		limit = 10
	}

	// Фильтрация
	title := c.DefaultQuery("title", "")
	query := database.DB

	if title != "" {
		query = query.Where("name ILIKE ?", "%"+title+"%")
	}

	// Применяем пагинацию
	query = query.Offset((page - 1) * limit).Limit(limit)

	if err := query.Find(&sneakers).Error; err != nil {
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный формат данных"})
		return
	}

	// Валидация данных
	if sneaker.Name == "" || sneaker.Price <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Название и цена кроссовки обязательны"})
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный формат данных"})
		return
	}

	if err := database.DB.Save(&sneaker).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при обновлении записи"})
		return
	}
	c.JSON(http.StatusOK, sneaker)
}

func DeleteSneaker(c *gin.Context) {
	result := database.DB.Delete(&models.Sneaker{}, c.Param("id"))
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при удалении кроссовки"})
		return
	}
	if result.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Кроссовка не найдена"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Кроссовка удалена"})
}


func SearchSneakers(c *gin.Context) {
	var sneakers []models.Sneaker
	category := c.DefaultQuery("category", "")
	brand := c.DefaultQuery("brand", "")
	priceMin := c.DefaultQuery("price_min", "0")
	priceMax := c.DefaultQuery("price_max", "10000")

	// Фильтрация
	query := database.DB.Where("category LIKE ?", "%"+category+"%").
		Where("brand LIKE ?", "%"+brand+"%").
		Where("price BETWEEN ? AND ?", priceMin, priceMax)

	if err := query.Find(&sneakers).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, sneakers)
}