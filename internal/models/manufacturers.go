package models

import (
	"time"
)

const TableNameManufacturer = "manufacturers"

// Manufacturer mapped from table <manufacturers>
type Manufacturer struct {
	ManufacturerID int32     `gorm:"column:manufacturer_id;primaryKey;autoIncrement:true" json:"manufacturer_id"`
	Name           string    `gorm:"column:name;not null" json:"name"`
	Country        string    `gorm:"column:country" json:"country"`
	Website        string    `gorm:"column:website" json:"website"`
	CreatedAt      time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
}

// TableName Manufacturer's table name
func (*Manufacturer) TableName() string {
	return TableNameManufacturer
}
