package main

import (
	"WeMarketOnGolang/internal/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Подключение к базе данных PostgreSQL
	dsn := "host=localhost user=postgres password=postgres dbname=wemarket port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Не удалось подключиться к базе данных: %v", err)
	}

	// Миграция всех моделей, возвращаемых функцией GetModels
	for _, model := range models.GetModels() {
		err := db.AutoMigrate(model)
		if err != nil {
			log.Fatalf("Не удалось выполнить миграцию для модели: %v", err)
		}
	}

	log.Println("Миграция выполнена успешно для всех моделей")
}
