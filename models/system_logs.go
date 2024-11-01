package models

import (
	"time"
)

const TableNameSystemLog = "system_logs"

// SystemLog mapped from table <system_logs>
type SystemLog struct {
	ID      int32     `gorm:"column:id;primaryKey;autoIncrement:true;comment:ID лога" json:"id"`                            // ID лога
	LogDate time.Time `gorm:"column:log_date;default:CURRENT_TIMESTAMP;comment:Дата и время записи лога" json:"log_date"`   // Дата и время записи лога
	Level   string    `gorm:"column:level;not null;comment:Уровень лога (информация, предупреждение, ошибка)" json:"level"` // Уровень лога (информация, предупреждение, ошибка)
	Message string    `gorm:"column:message;comment:Текст сообщения лога" json:"message"`                                   // Текст сообщения лога
	UserID  int32     `gorm:"column:user_id;comment:ID пользователя, связанного с логом" json:"user_id"`                    // ID пользователя, связанного с логом
}

// TableName SystemLog's table name
func (*SystemLog) TableName() string {
	return TableNameSystemLog
}
