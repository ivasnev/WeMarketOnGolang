package models

const TableNamePromotionStatus = "promotion_statuses"

// PromotionStatus mapped from table <promotion_statuses>
type PromotionStatus struct {
	ID   int32  `gorm:"column:id;type:integer;primaryKey;autoIncrement:true;comment:ID статуса" json:"id"`    // ID статуса
	Name string `gorm:"column:name;type:character varying(50);not null;comment:Название статуса" json:"name"` // Название статуса
}

// TableName PromotionStatus's table name
func (*PromotionStatus) TableName() string {
	return TableNamePromotionStatus
}
