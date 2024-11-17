package models

import (
	"time"
)

const TableNameNotification = "notifications"

// Notification mapped from table <notifications>
type Notification struct {
	ID               int32      `gorm:"column:id;type:integer;primaryKey;autoIncrement:true;comment:ID уведомления" json:"id"`                                                   // ID уведомления
	UserID           int32      `gorm:"column:user_id;type:integer;not null;comment:ID пользователя, которому отправлено уведомление" json:"user_id"`                            // ID пользователя, которому отправлено уведомление
	NotificationType string     `gorm:"column:notification_type;type:character varying(50);not null;comment:Тип уведомления (заказ, акция, системное)" json:"notification_type"` // Тип уведомления (заказ, акция, системное)
	Message          *string    `gorm:"column:message;type:text;comment:Текст уведомления" json:"message"`                                                                       // Текст уведомления
	SentDate         *time.Time `gorm:"column:sent_date;type:timestamp without time zone;default:CURRENT_TIMESTAMP;comment:Дата отправки уведомления" json:"sent_date"`          // Дата отправки уведомления
	ReadStatus       *bool      `gorm:"column:read_status;type:boolean;comment:Статус прочтения уведомления" json:"read_status"`                                                 // Статус прочтения уведомления
}

// TableName Notification's table name
func (*Notification) TableName() string {
	return TableNameNotification
}
