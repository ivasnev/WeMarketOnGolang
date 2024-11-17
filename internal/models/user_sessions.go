package models

import (
	"time"
)

const TableNameUserSession = "user_sessions"

// UserSession mapped from table <user_sessions>
type UserSession struct {
	ID        int32      `gorm:"column:id;type:integer;primaryKey;autoIncrement:true;comment:ID сессии" json:"id"`                                          // ID сессии
	UserID    int32      `gorm:"column:user_id;type:integer;not null;comment:ID пользователя" json:"user_id"`                                               // ID пользователя
	StartDate *time.Time `gorm:"column:start_date;type:timestamp without time zone;default:CURRENT_TIMESTAMP;comment:Дата начала сессии" json:"start_date"` // Дата начала сессии
	EndDate   *time.Time `gorm:"column:end_date;type:timestamp without time zone;comment:Дата окончания сессии" json:"end_date"`                            // Дата окончания сессии
	IPAddress *string    `gorm:"column:ip_address;type:character varying(45);comment:IP-адрес пользователя" json:"ip_address"`                              // IP-адрес пользователя
}

// TableName UserSession's table name
func (*UserSession) TableName() string {
	return TableNameUserSession
}
