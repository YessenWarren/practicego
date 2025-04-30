package handlers

import (
	"net/http"
	"sneakers/database"
	"sneakers/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Получение профиля текущего пользователя
func GetProfile(c *gin.Context) {
	// Получаем user_id из контекста, который был добавлен в middleware (после аутентификации)
	userID := c.GetUint("user_id")

	var user models.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Пользователь не найден"})
		return
	}

	// Возвращаем профиль пользователя без пароля
	c.JSON(http.StatusOK, gin.H{
		"id":       user.ID,
		"username": user.Username,
	})
}

// Обновление данных пользователя
func UpdateProfile(c *gin.Context) {
	// Получаем user_id из контекста
	userID := c.GetUint("user_id")

	var user models.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Пользователь не найден"})
		return
	}

	// Привязываем данные из запроса
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректные данные"})
		return
	}

	// Если пароль был обновлен, хэшируем его
	if input.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка шифрования пароля"})
			return
		}
		user.Password = string(hashedPassword)
	}

	// Обновляем данные пользователя
	user.Username = input.Username // Обновляем только имя пользователя
	// Если пароль был обновлен, он уже сохранен в переменной user.Password

	if err := database.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при обновлении профиля"})
		return
	}

	// Возвращаем обновленные данные без пароля
	c.JSON(http.StatusOK, gin.H{
		"id":       user.ID,
		"username": user.Username,
	})
}
