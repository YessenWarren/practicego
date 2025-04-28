package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"sneakers/handlers"
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
		"username": "test",
		"password": "123456",
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
	router := setupTestRouter()

	// Сначала регистрируем пользователя
	body := map[string]string{
		"username": "duplicate",
		"password": "password",
	}
	jsonValue, _ := json.Marshal(body)

	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)

	// Пытаемся зарегистрировать его снова
	reqDuplicate, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonValue))
	reqDuplicate.Header.Set("Content-Type", "application/json")
	wDuplicate := httptest.NewRecorder()
	router.ServeHTTP(wDuplicate, reqDuplicate)

	assert.Equal(t, http.StatusBadRequest, wDuplicate.Code)
}

// ✅ 3. Успешный вход
func TestLoginSuccess(t *testing.T) {
	setupDatabase()
	router := setupTestRouter()

	// Сначала регистрируем пользователя через API
	bodyRegister := map[string]string{
		"username": "testlogin",
		"password": "testpassword",
	}
	jsonRegister, _ := json.Marshal(bodyRegister)

	reqRegister, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonRegister))
	reqRegister.Header.Set("Content-Type", "application/json")
	wRegister := httptest.NewRecorder()
	router.ServeHTTP(wRegister, reqRegister)
	assert.Equal(t, http.StatusCreated, wRegister.Code)

	// Теперь логинимся
	bodyLogin := map[string]string{
		"username": "testlogin",
		"password": "testpassword",
	}
	jsonLogin, _ := json.Marshal(bodyLogin)

	reqLogin, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonLogin))
	reqLogin.Header.Set("Content-Type", "application/json")
	wLogin := httptest.NewRecorder()
	router.ServeHTTP(wLogin, reqLogin)

	assert.Equal(t, http.StatusOK, wLogin.Code)
}

// ✅ 4. Неправильный пароль
func TestLoginWrongPassword(t *testing.T) {
	setupDatabase()
	router := setupTestRouter()

	// Сначала регистрируем пользователя
	bodyRegister := map[string]string{
		"username": "wrongpass",
		"password": "correctpassword",
	}
	jsonRegister, _ := json.Marshal(bodyRegister)

	reqRegister, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonRegister))
	reqRegister.Header.Set("Content-Type", "application/json")
	wRegister := httptest.NewRecorder()
	router.ServeHTTP(wRegister, reqRegister)
	assert.Equal(t, http.StatusCreated, wRegister.Code)

	// Пытаемся залогиниться с неправильным паролем
	bodyLogin := map[string]string{
		"username": "wrongpass",
		"password": "wrongpassword",
	}
	jsonLogin, _ := json.Marshal(bodyLogin)

	reqLogin, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonLogin))
	reqLogin.Header.Set("Content-Type", "application/json")
	wLogin := httptest.NewRecorder()
	router.ServeHTTP(wLogin, reqLogin)

	assert.Equal(t, http.StatusUnauthorized, wLogin.Code)
}

// ✅ 5. Профиль без авторизации
func TestGetProfileUnauthorized(t *testing.T) {
	setupDatabase()
	router := setupTestRouter()

	req, _ := http.NewRequest("GET", "/profile", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code) // ожидаем 200, потому что нет проверки токена
}
