package models

import (
	"time"
)

const TableNameNotification = "notifications"

// Notification mapped from table <notifications>
type Notification struct {
	ID               int32     `gorm:"column:id;primaryKey;autoIncrement:true;comment:ID уведомления" json:"id"`                                     // ID уведомления
	UserID           int32     `gorm:"column:user_id;not null;comment:ID пользователя, которому отправлено уведомление" json:"user_id"`              // ID пользователя, которому отправлено уведомление
	NotificationType string    `gorm:"column:notification_type;not null;comment:Тип уведомления (заказ, акция, системное)" json:"notification_type"` // Тип уведомления (заказ, акция, системное)
	Message          string    `gorm:"column:message;comment:Текст уведомления" json:"message"`                                                      // Текст уведомления
	SentDate         time.Time `gorm:"column:sent_date;default:CURRENT_TIMESTAMP;comment:Дата отправки уведомления" json:"sent_date"`                // Дата отправки уведомления
	ReadStatus       bool      `gorm:"column:read_status;comment:Статус прочтения уведомления" json:"read_status"`                                   // Статус прочтения уведомления
}

// TableName Notification's table name
func (*Notification) TableName() string {
	return TableNameNotification
}
