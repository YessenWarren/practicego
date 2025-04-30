package handlers

import (
	"net/http"
	"sneakers/database"
	"sneakers/models"
	"github.com/gin-gonic/gin"
)

// Получение списка заказов с пагинацией
func GetOrders(c *gin.Context) {
	var orders []models.Order
	if err := database.DB.Find(&orders).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, orders)
}

// Создание нового заказа
func CreateOrder(c *gin.Context) {
	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректные данные"})
		return
	}

	// Проверка, существует ли пользователь
	var user models.User
	if err := database.DB.First(&user, order.UserID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Пользователь не найден"})
		return
	}

	// Проверка, существует ли кроссовка
	var sneaker models.Sneaker
	if err := database.DB.First(&sneaker, order.SneakerID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Кроссовки не найдены"})
		return
	}

	// Сохранение нового заказа в базе данных
	if err := database.DB.Create(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при создании заказа"})
		return
	}

	c.JSON(http.StatusCreated, order)
}

// Получение заказа по ID
func GetOrderByID(c *gin.Context) {
	var order models.Order
	if err := database.DB.First(&order, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Заказ не найден"})
		return
	}
	c.JSON(http.StatusOK, order)
}

// Обновление заказа
func UpdateOrder(c *gin.Context) {
	var order models.Order
	if err := database.DB.First(&order, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Заказ не найден"})
		return
	}

	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректные данные"})
		return
	}

	if err := database.DB.Save(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при обновлении заказа"})
		return
	}

	c.JSON(http.StatusOK, order)
}

// Удаление заказа
func DeleteOrder(c *gin.Context) {
	result := database.DB.Delete(&models.Order{}, c.Param("id"))
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при удалении заказа"})
		return
	}
	if result.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Заказ не найден"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Заказ удалён"})
}
