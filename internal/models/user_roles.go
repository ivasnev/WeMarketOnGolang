package models

const TableNameUserRole = "user_roles"

// UserRole mapped from table <user_roles>
type UserRole struct {
	ID   int32  `gorm:"column:id;primaryKey;autoIncrement:true;comment:ID роли" json:"id"` // ID роли
	Name string `gorm:"column:name;not null;comment:Название роли" json:"name"`            // Название роли
}

// TableName UserRole's table name
func (*UserRole) TableName() string {
	return TableNameUserRole
}
