package models

const TableNameReturnStatus = "return_statuses"

// ReturnStatus mapped from table <return_statuses>
type ReturnStatus struct {
	ID   int32  `gorm:"column:id;type:integer;primaryKey;autoIncrement:true;comment:ID статуса" json:"id"`    // ID статуса
	Name string `gorm:"column:name;type:character varying(50);not null;comment:Название статуса" json:"name"` // Название статуса
}

// TableName ReturnStatus's table name
func (*ReturnStatus) TableName() string {
	return TableNameReturnStatus
}
