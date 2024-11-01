package models

import (
	"time"
)

const TableNameSale = "sales"

// Sale mapped from table <sales>
type Sale struct {
	ID          int32     `gorm:"column:id;primaryKey;autoIncrement:true;comment:ID акции" json:"id"`               // ID акции
	Name        string    `gorm:"column:name;not null;comment:Название акции" json:"name"`                          // Название акции
	Description string    `gorm:"column:description;comment:Описание акции" json:"description"`                     // Описание акции
	Discount    float64   `gorm:"column:discount;not null;comment:Скидка по акции" json:"discount"`                 // Скидка по акции
	StartDate   time.Time `gorm:"column:start_date;comment:Дата начала акции" json:"start_date"`                    // Дата начала акции
	EndDate     time.Time `gorm:"column:end_date;comment:Дата окончания акции" json:"end_date"`                     // Дата окончания акции
	Products    string    `gorm:"column:products;comment:Список ID продуктов, участвующих в акции" json:"products"` // Список ID продуктов, участвующих в акции
}

// TableName Sale's table name
func (*Sale) TableName() string {
	return TableNameSale
}
