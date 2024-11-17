package models

import (
	"time"
)

const TableNameProduct = "products"

// Product mapped from table <products>
type Product struct {
	ID                 int32      `gorm:"column:id;type:integer;primaryKey;autoIncrement:true;comment:ID продукта" json:"id"`                                              // ID продукта
	Name               string     `gorm:"column:name;type:character varying(255);not null;comment:Название продукта" json:"name"`                                          // Название продукта
	Description        *string    `gorm:"column:description;type:text;comment:Описание продукта" json:"description"`                                                       // Описание продукта
	Price              float64    `gorm:"column:price;type:numeric(10,2);not null;comment:Цена продукта" json:"price"`                                                     // Цена продукта
	CategoryID         int32      `gorm:"column:category_id;type:integer;not null;comment:ID категории продукта" json:"category_id"`                                       // ID категории продукта
	Stock              int32      `gorm:"column:stock;type:integer;not null;comment:Количество на складе" json:"stock"`                                                    // Количество на складе
	ImageURL           *string    `gorm:"column:image_url;type:character varying(255);comment:URL изображения продукта" json:"image_url"`                                  // URL изображения продукта
	AddedDate          *time.Time `gorm:"column:added_date;type:timestamp without time zone;default:CURRENT_TIMESTAMP;comment:Дата добавления продукта" json:"added_date"` // Дата добавления продукта
	Sku                *string    `gorm:"column:sku;type:character varying(100);comment:Артикул продукта" json:"sku"`                                                      // Артикул продукта
	Weight             *float64   `gorm:"column:weight;type:numeric(10,2);comment:Вес продукта" json:"weight"`                                                             // Вес продукта
	AvailabilityStatus int32      `gorm:"column:availability_status;type:integer;not null;comment:ID статуса наличия продукта" json:"availability_status"`                 // ID статуса наличия продукта
	ManufacturerID     *int32     `gorm:"column:manufacturer_id;type:integer;comment:Производитель продукта" json:"manufacturer_id"`                                       // Производитель продукта
}

// TableName Product's table name
func (*Product) TableName() string {
	return TableNameProduct
}
