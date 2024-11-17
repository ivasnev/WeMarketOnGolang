package models

import (
	"time"
)

const TableNameSale = "sales"

// Sale mapped from table <sales>
type Sale struct {
	ID          int32      `gorm:"column:id;type:integer;primaryKey;autoIncrement:true;comment:ID акции" json:"id"`                // ID акции
	Name        string     `gorm:"column:name;type:character varying(100);not null;comment:Название акции" json:"name"`            // Название акции
	Description *string    `gorm:"column:description;type:text;comment:Описание акции" json:"description"`                         // Описание акции
	Discount    float64    `gorm:"column:discount;type:numeric(5,2);not null;comment:Скидка по акции" json:"discount"`             // Скидка по акции
	StartDate   *time.Time `gorm:"column:start_date;type:timestamp without time zone;comment:Дата начала акции" json:"start_date"` // Дата начала акции
	EndDate     *time.Time `gorm:"column:end_date;type:timestamp without time zone;comment:Дата окончания акции" json:"end_date"`  // Дата окончания акции
	Products    *string    `gorm:"column:products;type:jsonb;comment:Список ID продуктов, участвующих в акции" json:"products"`    // Список ID продуктов, участвующих в акции
}

// TableName Sale's table name
func (*Sale) TableName() string {
	return TableNameSale
}
