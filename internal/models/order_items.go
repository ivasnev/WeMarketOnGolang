package models

const TableNameOrderItem = "order_items"

// OrderItem mapped from table <order_items>
type OrderItem struct {
	ID        int32   `gorm:"column:id;type:integer;primaryKey;autoIncrement:true;comment:ID записи" json:"id"`     // ID записи
	OrderID   int32   `gorm:"column:order_id;type:integer;not null;comment:ID заказа" json:"order_id"`              // ID заказа
	ProductID int32   `gorm:"column:product_id;type:integer;not null;comment:ID продукта" json:"product_id"`        // ID продукта
	Quantity  int32   `gorm:"column:quantity;type:integer;not null;comment:Количество товара" json:"quantity"`      // Количество товара
	Price     float64 `gorm:"column:price;type:numeric(10,2);not null;comment:Цена за единицу товара" json:"price"` // Цена за единицу товара
}

// TableName OrderItem's table name
func (*OrderItem) TableName() string {
	return TableNameOrderItem
}
