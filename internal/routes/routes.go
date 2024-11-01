package routes

import (
	"WeMarketOnGolang/internal/handlers"
	"WeMarketOnGolang/internal/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRoutes(router *gin.Engine, db *gorm.DB) {
	productService := services.NewProductService(db)
	userService := services.NewUserService(db)
	productHandler := handlers.NewProductHandler(productService)
	userHandler := handlers.NewUserHandler(userService)

	api := router.Group("/v1")
	{
		// Группа маршрутов для продуктов
		products := api.Group("/products")
		{
			products.GET("/:id", productHandler.GetProduct)
			//products.POST("/", productHandler.CreateProduct)
			//products.PUT("/:id", productHandler.UpdateProduct)
			//products.DELETE("/:id", productHandler.DeleteProduct)
		}

		// Группа маршрутов для пользователей
		users := api.Group("/users")
		{
			users.GET("/:id", userHandler.GetUser)
			//users.POST("/", userHandler.CreateUser)
			//users.PUT("/:id", userHandler.UpdateUser)
			//users.DELETE("/:id", userHandler.DeleteUser)
		}
	}
}
