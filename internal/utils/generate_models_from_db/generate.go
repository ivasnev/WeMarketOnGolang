package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	// Настраиваем подключение к базе данных
	dsn := "host=localhost user=postgres password=postgres dbname=wemarket port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Создаем объект генератора
	g := gen.NewGenerator(gen.Config{
		FieldNullable:     true,
		FieldCoverable:    true,
		FieldSignable:     true,
		FieldWithIndexTag: true,
		FieldWithTypeTag:  true,
		OutPath:           "./internal/models/generated", // Путь для сохранения сгенерированных моделей
		ModelPkgPath:      "internal/models",
		Mode:              gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	})

	// Используем подключение к базе данных
	g.UseDB(db)

	// Автоматически сгенерировать модели на основе существующих таблиц
	g.GenerateAllTable()

	// Выполняем генерацию
	g.Execute()
}
