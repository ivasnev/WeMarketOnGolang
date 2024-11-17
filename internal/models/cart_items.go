package models

const TableNameCartItem = "cart_items"

// CartItem mapped from table <cart_items>
type CartItem struct {
	ID        int32 `gorm:"column:id;type:integer;primaryKey;autoIncrement:true;comment:ID записи" json:"id"` // ID записи
	CartID    int32 `gorm:"column:cart_id;type:integer;not null;comment:ID корзины" json:"cart_id"`           // ID корзины
	ProductID int32 `gorm:"column:product_id;type:integer;not null;comment:ID продукта" json:"product_id"`    // ID продукта
	Quantity  int32 `gorm:"column:quantity;type:integer;not null;comment:Количество товара" json:"quantity"`  // Количество товара
}

// TableName CartItem's table name
func (*CartItem) TableName() string {
	return TableNameCartItem
}
