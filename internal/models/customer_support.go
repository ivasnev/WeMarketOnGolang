package models

import (
	"time"
)

const TableNameCustomerSupport = "customer_support"

// CustomerSupport mapped from table <customer_support>
type CustomerSupport struct {
	ID          int32      `gorm:"column:id;type:integer;primaryKey;autoIncrement:true;comment:ID обращения" json:"id"`                                                // ID обращения
	UserID      int32      `gorm:"column:user_id;type:integer;not null;comment:ID пользователя, сделавшего обращение" json:"user_id"`                                  // ID пользователя, сделавшего обращение
	Subject     *string    `gorm:"column:subject;type:character varying(255);comment:Тема обращения" json:"subject"`                                                   // Тема обращения
	Message     *string    `gorm:"column:message;type:text;comment:Текст обращения" json:"message"`                                                                    // Текст обращения
	StatusID    int32      `gorm:"column:status_id;type:integer;not null;comment:ID статуса обращения" json:"status_id"`                                               // ID статуса обращения
	CreatedDate *time.Time `gorm:"column:created_date;type:timestamp without time zone;default:CURRENT_TIMESTAMP;comment:Дата создания обращения" json:"created_date"` // Дата создания обращения
	Response    *string    `gorm:"column:response;type:text;comment:Ответ сотрудника поддержки" json:"response"`                                                       // Ответ сотрудника поддержки
}

// TableName CustomerSupport's table name
func (*CustomerSupport) TableName() string {
	return TableNameCustomerSupport
}
