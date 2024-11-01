package models

import (
	"time"
)

const TableNamePromotion = "promotions"

// Promotion mapped from table <promotions>
type Promotion struct {
	ID                 int32     `gorm:"column:id;primaryKey;autoIncrement:true;comment:ID промокода" json:"id"`                                               // ID промокода
	Code               string    `gorm:"column:code;not null;comment:Код промокода" json:"code"`                                                               // Код промокода
	Discount           float64   `gorm:"column:discount;not null;comment:Скидка, предоставляемая промокодом" json:"discount"`                                  // Скидка, предоставляемая промокодом
	MinimumOrderAmount float64   `gorm:"column:minimum_order_amount;comment:Минимальная сумма заказа для использования промокода" json:"minimum_order_amount"` // Минимальная сумма заказа для использования промокода
	StartDate          time.Time `gorm:"column:start_date;comment:Дата начала действия промокода" json:"start_date"`                                           // Дата начала действия промокода
	EndDate            time.Time `gorm:"column:end_date;comment:Дата окончания действия промокода" json:"end_date"`                                            // Дата окончания действия промокода
	StatusID           int32     `gorm:"column:status_id;not null;comment:ID статуса промокода" json:"status_id"`                                              // ID статуса промокода
	UsageLimit         int32     `gorm:"column:usage_limit;comment:Лимит использования промокода" json:"usage_limit"`                                          // Лимит использования промокода
}

// TableName Promotion's table name
func (*Promotion) TableName() string {
	return TableNamePromotion
}
