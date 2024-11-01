package main

import (
	"WeMarketOnGolang/internal/routes"
	"WeMarketOnGolang/pkg"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func main() {
	// Инициализация логгера
	pkg.InitLogger()

	// Загрузка конфигурации
	config, err := pkg.LoadConfig("configs/config.yaml")
	if err != nil {
		pkg.Error("Не удалось загрузить конфигурацию")
		return
	}

	databaseURL := pkg.GetDBUrl(config)

	// Инициализация базы данных
	db, err := gorm.Open(postgres.Open(databaseURL), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Миграция моделей
	//if err := db.AutoMigrate(&models.User{}, &models.Product{}); err != nil {
	//	log.Fatalf("Failed to migrate models: %v", err)
	//}

	// Создание нового роутера Gin
	router := gin.Default()

	// Установка маршрутов
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Сервер работает на порту " + config.Server.Port,
		})
	})

	routes.InitRoutes(router, db)

	// Запуск сервера
	pkg.Info("Запуск сервера на порту " + config.Server.Port)
	if err := router.Run(":" + config.Server.Port); err != nil {
		pkg.Error("Ошибка при запуске сервера: " + err.Error())
	}
}
