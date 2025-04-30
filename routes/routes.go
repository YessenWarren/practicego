package routes

import (
	"github.com/gin-gonic/gin"
	"sneakers/handlers"
	"sneakers/middleware"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)

	api := r.Group("/api")
	api.Use(middleware.AuthMiddleware())
	{
		// 1. Работа с кроссовками
		api.GET("/sneakers", handlers.GetSneakers)        // Список кроссовок с пагинацией и фильтрацией
		api.POST("/sneakers", handlers.CreateSneaker)
		api.PUT("/sneakers/:id", handlers.UpdateSneaker)
		api.DELETE("/sneakers/:id", handlers.DeleteSneaker)
		api.GET("/sneakers/:id", handlers.GetSneakerByID)

		// 2. Работа с категориями
		api.GET("/categories", handlers.GetCategories)   // Список категорий с пагинацией
		api.POST("/categories", handlers.CreateCategory)

		// 3. Работа с брендами
		api.GET("/brands", handlers.GetBrands)           // Список брендов с пагинацией
		api.POST("/brands", handlers.CreateBrand)

		// 4. Работа с пользователями (профиль)
		api.GET("/profile", handlers.GetProfile)         // Получение профиля пользователя
		api.PUT("/profile", handlers.UpdateProfile)      // Обновление профиля пользователя

		// 5. Работа с заказами
		api.GET("/orders", handlers.GetOrders)           // Список заказов с пагинацией
		api.POST("/orders", handlers.CreateOrder)        // Создание заказа
		api.GET("/orders/:id", handlers.GetOrderByID)    // Получение заказа по ID
		api.PUT("/orders/:id", handlers.UpdateOrder)     // Обновление заказа
		api.DELETE("/orders/:id", handlers.DeleteOrder)  // Удаление заказа

		// 6. Работа с отзывами
		api.GET("/reviews", handlers.GetReviews)         // Список отзывов с пагинацией
		api.POST("/reviews", handlers.CreateReview)      // Создание отзыва
		api.GET("/reviews/:id", handlers.GetReviewByID)  // Получение отзыва по ID
		api.PUT("/reviews/:id", handlers.UpdateReview)   // Обновление отзыва
		api.DELETE("/reviews/:id", handlers.DeleteReview)// Удаление отзыва

		// 7. Статистика
		api.GET("/statistics/sales", handlers.GetSalesStatistics) // Статистика по продажам
		api.GET("/statistics/users", handlers.GetUsersStatistics) // Статистика по пользователям

		// 8. Дополнительные маршруты для фильтрации
		api.GET("/sneakers/search", handlers.SearchSneakers) // Поиск кроссовок по фильтрам
	}
}
