package models

import (
	"time"
)

const TableNameProductChange = "product_changes"

// ProductChange mapped from table <product_changes>
type ProductChange struct {
	ID           int32     `gorm:"column:id;primaryKey;autoIncrement:true;comment:ID изменения" json:"id"`                 // ID изменения
	ProductID    int32     `gorm:"column:product_id;not null;comment:ID продукта" json:"product_id"`                       // ID продукта
	FieldChanged string    `gorm:"column:field_changed;not null;comment:Поле, которое изменилось" json:"field_changed"`    // Поле, которое изменилось
	OldValue     string    `gorm:"column:old_value;comment:Старое значение" json:"old_value"`                              // Старое значение
	NewValue     string    `gorm:"column:new_value;comment:Новое значение" json:"new_value"`                               // Новое значение
	ChangeDate   time.Time `gorm:"column:change_date;default:CURRENT_TIMESTAMP;comment:Дата изменения" json:"change_date"` // Дата изменения
	ChangedBy    int32     `gorm:"column:changed_by;comment:ID пользователя, который внёс изменение" json:"changed_by"`    // ID пользователя, который внёс изменение
}

// TableName ProductChange's table name
func (*ProductChange) TableName() string {
	return TableNameProductChange
}
