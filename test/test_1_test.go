package test

import (
	"bytes"
	"sneakers/handlers"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// Вспомогательная функция для настройки маршрутов
func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/sneakers", handlers.GetSneakers)
    r.GET("/sneakers/:id", handlers.GetSneakerByID)
    r.POST("/sneakers", handlers.CreateSneaker)
    r.PUT("/sneakers/:id", handlers.UpdateSneaker)
    r.DELETE("/sneakers/:id", handlers.DeleteSneaker)
	return r
}

// ✅ 1. Тест на получение всех кроссовок
func TestGetSneakers(t *testing.T) {
	// Инициализация базы данных перед тестом
	setupDatabase()

	// Настройка маршрутов
	router := setupRouter()

	// Тестовый запрос
	req, _ := http.NewRequest("GET", "/sneakers", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Проверка ответа
	assert.Equal(t, http.StatusOK, resp.Code)
}

// ✅ 2. Тест на получение кроссовки по ID (не найдена)
func TestGetSneakerByID_NotFound(t *testing.T) {
	// Инициализация базы данных перед тестом
	setupDatabase()

	// Настройка маршрутов
	router := setupRouter()

	// Тестовый запрос
	req, _ := http.NewRequest("GET", "/sneakers/9999", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Проверка ответа
	assert.Equal(t, http.StatusNotFound, resp.Code)
}

// ✅ 3. Тест на создание кроссовки с некорректным JSON
func TestCreateSneaker_BadRequest(t *testing.T) {
	// Инициализация базы данных перед тестом
	setupDatabase()

	// Настройка маршрутов
	router := setupRouter()

	// Тестовые данные с ошибочным полем
	reqBody := `{"bad_field":"value"}`
	req, _ := http.NewRequest("POST", "/sneakers", bytes.NewBufferString(reqBody))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Проверка ответа
	assert.Equal(t, http.StatusBadRequest, resp.Code)
}

// ✅ 4. Тест на обновление несуществующей кроссовки
func TestUpdateSneaker_NotFound(t *testing.T) {
	// Инициализация базы данных перед тестом
	setupDatabase()

	// Настройка маршрутов
	router := setupRouter()

	// Тестовые данные для обновления
	reqBody := `{"name":"Updated Sneaker"}`
	req, _ := http.NewRequest("PUT", "/sneakers/9999", bytes.NewBufferString(reqBody))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Проверка ответа
	assert.Equal(t, http.StatusNotFound, resp.Code)
}

// ✅ 5. Тест на удаление несуществующей кроссовки
func TestDeleteSneaker_BadRequest(t *testing.T) {
	// Инициализация базы данных перед тестом
	setupDatabase()

	// Настройка маршрутов
	router := setupRouter()

	// Тестовый запрос на удаление
	req, _ := http.NewRequest("DELETE", "/sneakers/9999", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Проверка ответа
	assert.Equal(t, http.StatusBadRequest, resp.Code)
}
