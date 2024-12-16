package main

import (
	"WeMarketOnGolang/docs"
	"WeMarketOnGolang/internal/routes"
	"WeMarketOnGolang/internal/utils"
	"WeMarketOnGolang/pkg"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

// @title           WeMarket API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/
// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io
// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
// @host      localhost:8080
// @BasePath  /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @tokenUrl /v1/auth/jwt/login
// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {

	docs.SwaggerInfo.Host = utils.GetDynamicHost()

	// Инициализация логгера
	pkg.InitLogger()

	// Загрузка конфигурации
	config, err := pkg.LoadConfig("configs/config_local_dev.yaml")
	if err != nil {
		pkg.Error("Не удалось загрузить конфигурацию")
		return
	}

	databaseURL := pkg.GetDBUrl(config)

	newDBLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // Вывод логов в консоль
		logger.Config{
			SlowThreshold:             time.Second, // Порог "медленного" запроса
			LogLevel:                  logger.Info, // Уровень логирования
			IgnoreRecordNotFoundError: true,        // Игнорировать ошибки RecordNotFound
			Colorful:                  true,        // Цветные логи (для консоли)
		},
	)

	// Инициализация базы данных
	db, err := gorm.Open(postgres.Open(databaseURL), &gorm.Config{
		Logger: newDBLogger,
	})
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
