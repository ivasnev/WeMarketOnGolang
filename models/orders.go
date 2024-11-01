package models

import (
	"time"
)

const TableNameOrder = "orders"

// Order mapped from table <orders>
type Order struct {
	ID              int32     `gorm:"column:id;primaryKey;autoIncrement:true;comment:ID заказа" json:"id"`                  // ID заказа
	UserID          int32     `gorm:"column:user_id;not null;comment:ID пользователя, который сделал заказ" json:"user_id"` // ID пользователя, который сделал заказ
	OrderDate       time.Time `gorm:"column:order_date;default:CURRENT_TIMESTAMP;comment:Дата заказа" json:"order_date"`    // Дата заказа
	TotalAmount     float64   `gorm:"column:total_amount;not null;comment:Итоговая сумма заказа" json:"total_amount"`       // Итоговая сумма заказа
	StatusID        int32     `gorm:"column:status_id;not null;comment:ID статуса заказа" json:"status_id"`                 // ID статуса заказа
	DeliveryAddress string    `gorm:"column:delivery_address;comment:Адрес доставки" json:"delivery_address"`               // Адрес доставки
	PaymentMethod   string    `gorm:"column:payment_method;comment:Метод оплаты" json:"payment_method"`                     // Метод оплаты
	TrackingNumber  string    `gorm:"column:tracking_number;comment:Номер отслеживания заказа" json:"tracking_number"`      // Номер отслеживания заказа
	FulfillmentTime time.Time `gorm:"column:fulfillment_time;comment:Время выполнения заказа" json:"fulfillment_time"`      // Время выполнения заказа
	ShippingMethod  string    `gorm:"column:shipping_method;comment:Метод доставки" json:"shipping_method"`                 // Метод доставки
}

// TableName Order's table name
func (*Order) TableName() string {
	return TableNameOrder
}
