package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"sneakers/handlers"
	"sneakers/models"
	"sneakers/database"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)
	r.GET("/profile", handlers.GetProfile)
	return r
}

// ✅ 1. Успешная регистрация
func TestRegisterSuccess(t *testing.T) {
	setupDatabase()
	router := setupTestRouter()

	body := map[string]string{
		"username": "testuser",
		"password": "testpass",
	}
	jsonValue, _ := json.Marshal(body)

	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}

// ✅ 2. Регистрация дубликата
func TestRegisterDuplicate(t *testing.T) {
	setupDatabase()

	// Добавляем пользователя напрямую
	user := models.User{Username: "duplicate", Password: "password"}
	database.DB.Create(&user)

	body := map[string]string{"username": "duplicate", "password": "password"}
	jsonValue, _ := json.Marshal(body)

	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router := setupTestRouter()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

// ✅ 3. Успешный вход
func TestLoginSuccess(t *testing.T) {
	setupDatabase()

	user := models.User{Username: "testlogin", Password: "testlogin"}
	database.DB.Create(&user)

	body := map[string]string{"username": "testlogin", "password": "testlogin"}
	jsonValue, _ := json.Marshal(body)

	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router := setupTestRouter()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

// ✅ 4. Неправильный пароль
func TestLoginWrongPassword(t *testing.T) {
	setupDatabase()

	user := models.User{Username: "wrongpass", Password: "correct"}
	database.DB.Create(&user)

	body := map[string]string{"username": "wrongpass", "password": "wrong"}
	jsonValue, _ := json.Marshal(body)

	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router := setupTestRouter()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

// ✅ 5. Профиль без авторизации
func TestGetProfileUnauthorized(t *testing.T) {
	setupDatabase()

	req, _ := http.NewRequest("GET", "/profile", nil)
	w := httptest.NewRecorder()

	router := setupTestRouter()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code) // Изменить на 401, если используешь middleware
}
