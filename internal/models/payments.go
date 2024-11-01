package models

import (
	"time"
)

const TableNamePayment = "payments"

// Payment mapped from table <payments>
type Payment struct {
	ID            int32     `gorm:"column:id;primaryKey;autoIncrement:true;comment:ID платежа" json:"id"`                   // ID платежа
	OrderID       int32     `gorm:"column:order_id;not null;comment:ID заказа" json:"order_id"`                             // ID заказа
	PaymentMethod string    `gorm:"column:payment_method;not null;comment:Метод оплаты" json:"payment_method"`              // Метод оплаты
	StatusID      int32     `gorm:"column:status_id;not null;comment:ID статуса платежа" json:"status_id"`                  // ID статуса платежа
	PaymentDate   time.Time `gorm:"column:payment_date;default:CURRENT_TIMESTAMP;comment:Дата платежа" json:"payment_date"` // Дата платежа
	Amount        float64   `gorm:"column:amount;not null;comment:Сумма платежа" json:"amount"`                             // Сумма платежа
	TransactionID string    `gorm:"column:transaction_id;comment:Транзакционный идентификатор" json:"transaction_id"`       // Транзакционный идентификатор
}

// TableName Payment's table name
func (*Payment) TableName() string {
	return TableNamePayment
}
