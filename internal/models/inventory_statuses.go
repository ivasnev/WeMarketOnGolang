package models

const TableNameInventoryStatus = "inventory_statuses"

// InventoryStatus mapped from table <inventory_statuses>
type InventoryStatus struct {
	ID   int32  `gorm:"column:id;primaryKey;autoIncrement:true;comment:ID статуса" json:"id"` // ID статуса
	Name string `gorm:"column:name;not null;comment:Название статуса" json:"name"`            // Название статуса
}

// TableName InventoryStatus's table name
func (*InventoryStatus) TableName() string {
	return TableNameInventoryStatus
}
