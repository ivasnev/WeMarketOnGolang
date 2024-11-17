package models

import (
	"time"
)

const TableNamePayment = "payments"

// Payment mapped from table <payments>
type Payment struct {
	ID            int32      `gorm:"column:id;type:integer;primaryKey;autoIncrement:true;comment:ID платежа" json:"id"`                                       // ID платежа
	OrderID       int32      `gorm:"column:order_id;type:integer;not null;comment:ID заказа" json:"order_id"`                                                 // ID заказа
	PaymentMethod string     `gorm:"column:payment_method;type:character varying(50);not null;comment:Метод оплаты" json:"payment_method"`                    // Метод оплаты
	StatusID      int32      `gorm:"column:status_id;type:integer;not null;comment:ID статуса платежа" json:"status_id"`                                      // ID статуса платежа
	PaymentDate   *time.Time `gorm:"column:payment_date;type:timestamp without time zone;default:CURRENT_TIMESTAMP;comment:Дата платежа" json:"payment_date"` // Дата платежа
	Amount        float64    `gorm:"column:amount;type:numeric(10,2);not null;comment:Сумма платежа" json:"amount"`                                           // Сумма платежа
	TransactionID *string    `gorm:"column:transaction_id;type:character varying(100);comment:Транзакционный идентификатор" json:"transaction_id"`            // Транзакционный идентификатор
}

// TableName Payment's table name
func (*Payment) TableName() string {
	return TableNamePayment
}
