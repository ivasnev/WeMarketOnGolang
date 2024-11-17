package models

import (
	"time"
)

const TableNameManufacturer = "manufacturers"

// Manufacturer mapped from table <manufacturers>
type Manufacturer struct {
	ManufacturerID int32      `gorm:"column:manufacturer_id;type:integer;primaryKey;autoIncrement:true" json:"manufacturer_id"`
	Name           string     `gorm:"column:name;type:character varying(255);not null" json:"name"`
	Country        *string    `gorm:"column:country;type:character varying(100)" json:"country"`
	Website        *string    `gorm:"column:website;type:character varying(255)" json:"website"`
	CreatedAt      *time.Time `gorm:"column:created_at;type:timestamp without time zone;default:CURRENT_TIMESTAMP" json:"created_at"`
}

// TableName Manufacturer's table name
func (*Manufacturer) TableName() string {
	return TableNameManufacturer
}
