package routes

import (
	"WeMarketOnGolang/internal"
	"WeMarketOnGolang/internal/handlers"
	"WeMarketOnGolang/internal/middleware"
	"WeMarketOnGolang/internal/services"
	"WeMarketOnGolang/internal/services/categories"
	"WeMarketOnGolang/internal/services/inventoryStatus"
	"WeMarketOnGolang/internal/services/products"
	"WeMarketOnGolang/internal/services/tasks"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
)

func InitRoutes(router *gin.Engine, db *gorm.DB) {
	productService := products.NewProductService(db)
	productServiceV0 := products.NewInMemoryProductService()
	productServiceV0.SeedProducts()
	productHandler := handlers.NewProductHandler(productService)
	productHandlerV0 := handlers.NewProductHandler(productServiceV0)
	authService := services.NewJWTAuthService(internal.JWTSecretKey, db)
	authHandler := handlers.NewAuthHandler(authService)
	userService := services.NewUserService(db)
	userHandler := handlers.NewUserHandler(userService)
	categoryService := categories.NewCategoryService(db)
	categoryHandler := handlers.NewCategoryHandler(categoryService)
	inventoryStatusService := inventoryStatus.NewInventoryStatusService(db)
	inventoryStatusHandler := handlers.NewInventoryStatusHandler(inventoryStatusService)
	taskService := tasks.NewTaskService(5) // Создаем сервис задач
	taskHandler := handlers.NewTaskHandler(taskService)

	apiV0 := router.Group("/v0")
	{
		// Группа маршрутов для продуктов
		products := apiV0.Group("/products")
		{
			products.GET("/:id", productHandlerV0.GetProductByID)
			products.GET("/", productHandlerV0.GetAllProducts)
			products.POST("/", productHandlerV0.CreateProduct)
			products.PUT("/:id", productHandlerV0.UpdateProduct)
			products.DELETE("/:id", productHandlerV0.DeleteProduct)
		}
	}

	apiV1 := router.Group("/v1")
	{
		apiV1.Use(middleware.TimeoutMiddleware(5 * time.Second))
		// Группа маршрутов для продуктов
		products := apiV1.Group("/products")
		{
			products.Use(middleware.JWTMiddleware())
			products.GET("/:id", productHandler.GetProductByID)
			products.GET("/", productHandler.GetAllProducts)
			products.POST("/", productHandler.CreateProduct)
			products.PUT("/:id", productHandler.UpdateProduct)
			products.DELETE("/:id", productHandler.DeleteProduct)
		}
		categories := apiV1.Group("/category")
		{
			categories.Use(middleware.JWTMiddleware())
			categories.GET("/", categoryHandler.GetAllCategories)
			categories.POST("/", categoryHandler.CreateCategory)
			categories.GET("/:id", categoryHandler.GetCategory)
			categories.PUT("/:id", categoryHandler.UpdateCategory)
			categories.DELETE("/:id", categoryHandler.DeleteCategory)
		}
		inventoryStatuses := apiV1.Group("/inventory-statuses")
		{
			inventoryStatuses.Use(middleware.JWTMiddleware())
			inventoryStatuses.POST("/", inventoryStatusHandler.CreateInventoryStatus)
			inventoryStatuses.GET("/:id", inventoryStatusHandler.GetInventoryStatusByID)
			inventoryStatuses.GET("/", inventoryStatusHandler.GetAllInventoryStatuses)
			inventoryStatuses.PUT("/:id", inventoryStatusHandler.UpdateInventoryStatus)
			inventoryStatuses.DELETE("/:id", inventoryStatusHandler.DeleteInventoryStatus)
		}

		authGroup := apiV1.Group("/auth")
		{
			authGroup.POST("/jwt/login", authHandler.Login)
			authGroup.POST("/jwt/logout", middleware.JWTMiddleware(), authHandler.Logout)
			authGroup.POST("/register", userHandler.Register)
			//authGroup.POST("/forgot-password", userHandler.ForgotPassword)
			//authGroup.POST("/reset-password", userHandler.ResetPassword)
			//authGroup.POST("/request-verify-token", userHandler.RequestVerifyToken)
			//authGroup.POST("/verify", userHandler.Verify)
		}

		usersGroup := apiV1.Group("/users")
		{
			usersGroup.Use(middleware.JWTMiddleware())
			usersGroup.GET("/me", userHandler.GetCurrentUser)
			usersGroup.PATCH("/me", userHandler.UpdateCurrentUser)
			usersGroup.GET("/:id", userHandler.GetUserByID)
			//usersGroup.PATCH("/:id", userHandler.UpdateUserByID)
			//usersGroup.DELETE("/:id", userHandler.DeleteUserByID)
		}

		tasksGroup := apiV1.Group("/tasks")
		{
			// Управление задачами
			tasksGroup.POST("/inf", taskHandler.CreateTaskInf)
			tasksGroup.POST("/classic", taskHandler.CreateTaskClassic)
			tasksGroup.GET("", taskHandler.GetAllTasks)
			tasksGroup.GET("/:id", taskHandler.GetTask)
			tasksGroup.DELETE("/:id", taskHandler.CancelTask)
		}
	}
}
