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
		api.GET("/sneakers", handlers.GetSneakers)
		api.POST("/sneakers", handlers.CreateSneaker)
		api.PUT("/sneakers/:id", handlers.UpdateSneaker)
		api.DELETE("/sneakers/:id", handlers.DeleteSneaker)

		api.GET("/sneakers/:id", handlers.GetSneakerByID)
		api.GET("/categories", handlers.GetCategories)
		api.POST("/categories", handlers.CreateCategory)
		api.GET("/brands", handlers.GetBrands)
		api.POST("/brands", handlers.CreateBrand)

		api.GET("/profile", handlers.GetProfile)
	}
}