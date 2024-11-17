package routes

import (
	"WeMarketOnGolang/internal/handlers"
	"WeMarketOnGolang/internal/services"
	"WeMarketOnGolang/internal/services/categories"
	"WeMarketOnGolang/internal/services/products"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRoutes(router *gin.Engine, db *gorm.DB) {
	productService := products.NewProductService(db)
	productServiceV0 := products.NewInMemoryProductService()
	productServiceV0.SeedProducts()
	userService := services.NewUserService(db)
	productHandler := handlers.NewProductHandler(productService)
	productHandlerV0 := handlers.NewProductHandler(productServiceV0)
	userHandler := handlers.NewUserHandler(userService)
	categoryService := categories.NewCategoryService(db)
	categoryHandler := handlers.NewCategoryHandler(categoryService)

	apiV0 := router.Group("/v0")
	{
		// Группа маршрутов для продуктов
		products := apiV0.Group("/products")
		{
			products.GET("/:id", productHandlerV0.GetProduct)
			products.GET("/", productHandlerV0.ListProducts)
			products.POST("/", productHandlerV0.CreateProduct)
			products.PUT("/:id", productHandlerV0.UpdateProduct)
			products.DELETE("/:id", productHandlerV0.DeleteProduct)
		}
	}

	apiV1 := router.Group("/v1")
	{
		// Группа маршрутов для продуктов
		products := apiV1.Group("/products")
		{
			products.GET("/:id", productHandler.GetProduct)
			products.GET("/", productHandler.ListProducts)
			products.POST("/", productHandler.CreateProduct)
			products.PUT("/:id", productHandler.UpdateProduct)
			products.DELETE("/:id", productHandler.DeleteProduct)
		}
		categories := apiV1.Group("/category")
		{
			categories.GET("/", categoryHandler.GetAllCategories)
			categories.POST("/", categoryHandler.CreateCategory)
			categories.GET("/:id", categoryHandler.GetCategory)
			categories.PUT("/:id", categoryHandler.UpdateCategory)
			categories.DELETE("/:id", categoryHandler.DeleteCategory)
		}

		// Группа маршрутов для пользователей
		users := apiV1.Group("/users")
		{
			users.GET("/:id", userHandler.GetUser)
			//users.POST("/", userHandler.CreateUser)
			//users.PUT("/:id", userHandler.UpdateUser)
			//users.DELETE("/:id", userHandler.DeleteUser)
		}
	}
}
