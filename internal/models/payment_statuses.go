package models

const TableNamePaymentStatus = "payment_statuses"

// PaymentStatus mapped from table <payment_statuses>
type PaymentStatus struct {
	ID   int32  `gorm:"column:id;type:integer;primaryKey;autoIncrement:true;comment:ID статуса" json:"id"`    // ID статуса
	Name string `gorm:"column:name;type:character varying(50);not null;comment:Название статуса" json:"name"` // Название статуса
}

// TableName PaymentStatus's table name
func (*PaymentStatus) TableName() string {
	return TableNamePaymentStatus
}
