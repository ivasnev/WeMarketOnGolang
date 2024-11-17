package models

import (
	"time"
)

const TableNameReturn = "returns"

// Return mapped from table <returns>
type Return struct {
	ID           int32      `gorm:"column:id;type:integer;primaryKey;autoIncrement:true;comment:ID возврата" json:"id"`                                     // ID возврата
	OrderID      int32      `gorm:"column:order_id;type:integer;not null;comment:ID заказа" json:"order_id"`                                                // ID заказа
	ProductID    int32      `gorm:"column:product_id;type:integer;not null;comment:ID продукта" json:"product_id"`                                          // ID продукта
	ReturnDate   *time.Time `gorm:"column:return_date;type:timestamp without time zone;default:CURRENT_TIMESTAMP;comment:Дата возврата" json:"return_date"` // Дата возврата
	StatusID     int32      `gorm:"column:status_id;type:integer;not null;comment:ID статуса возврата" json:"status_id"`                                    // ID статуса возврата
	ReturnReason *string    `gorm:"column:return_reason;type:text;comment:Причина возврата" json:"return_reason"`                                           // Причина возврата
}

// TableName Return's table name
func (*Return) TableName() string {
	return TableNameReturn
}
