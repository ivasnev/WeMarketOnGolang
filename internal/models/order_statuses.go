package models

const TableNameOrderStatus = "order_statuses"

// OrderStatus mapped from table <order_statuses>
type OrderStatus struct {
	ID   int32  `gorm:"column:id;type:integer;primaryKey;autoIncrement:true;comment:ID статуса" json:"id"`    // ID статуса
	Name string `gorm:"column:name;type:character varying(50);not null;comment:Название статуса" json:"name"` // Название статуса
}

// TableName OrderStatus's table name
func (*OrderStatus) TableName() string {
	return TableNameOrderStatus
}
