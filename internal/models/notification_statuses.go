package models

const TableNameNotificationStatus = "notification_statuses"

// NotificationStatus mapped from table <notification_statuses>
type NotificationStatus struct {
	ID   int32  `gorm:"column:id;type:integer;primaryKey;autoIncrement:true;comment:ID статуса" json:"id"`    // ID статуса
	Name string `gorm:"column:name;type:character varying(50);not null;comment:Название статуса" json:"name"` // Название статуса
}

// TableName NotificationStatus's table name
func (*NotificationStatus) TableName() string {
	return TableNameNotificationStatus
}
