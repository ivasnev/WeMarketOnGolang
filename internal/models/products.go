package models

import (
	"time"
)

const TableNameProduct = "products"

// Product mapped from table <products>
type Product struct {
	ID                 int32     `gorm:"column:id;primaryKey;autoIncrement:true;comment:ID продукта" json:"id"`                              // ID продукта
	Name               string    `gorm:"column:name;not null;comment:Название продукта" json:"name"`                                         // Название продукта
	Description        string    `gorm:"column:description;comment:Описание продукта" json:"description"`                                    // Описание продукта
	Price              float64   `gorm:"column:price;not null;comment:Цена продукта" json:"price"`                                           // Цена продукта
	CategoryID         int32     `gorm:"column:category_id;not null;comment:ID категории продукта" json:"category_id"`                       // ID категории продукта
	Stock              int32     `gorm:"column:stock;not null;comment:Количество на складе" json:"stock"`                                    // Количество на складе
	ImageURL           string    `gorm:"column:image_url;comment:URL изображения продукта" json:"image_url"`                                 // URL изображения продукта
	AddedDate          time.Time `gorm:"column:added_date;default:CURRENT_TIMESTAMP;comment:Дата добавления продукта" json:"added_date"`     // Дата добавления продукта
	Options            string    `gorm:"column:options;comment:Варианты продукта (цвет, размер и т.д.)" json:"options"`                      // Варианты продукта (цвет, размер и т.д.)
	Sku                string    `gorm:"column:sku;comment:Артикул продукта" json:"sku"`                                                     // Артикул продукта
	Weight             float64   `gorm:"column:weight;comment:Вес продукта" json:"weight"`                                                   // Вес продукта
	Dimensions         string    `gorm:"column:dimensions;comment:Габариты продукта (длина, ширина, высота)" json:"dimensions"`              // Габариты продукта (длина, ширина, высота)
	AvailabilityStatus int32     `gorm:"column:availability_status;not null;comment:ID статуса наличия продукта" json:"availability_status"` // ID статуса наличия продукта
	ManufacturerID     int32     `gorm:"column:manufacturer_id;comment:Производитель продукта" json:"manufacturer_id"`                       // Производитель продукта
}

// TableName Product's table name
func (*Product) TableName() string {
	return TableNameProduct
}
