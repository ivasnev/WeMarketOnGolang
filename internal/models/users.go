package models

import (
	"time"
)

const TableNameUser = "users"

// User mapped from table <users>
type User struct {
	ID               int32      `gorm:"column:id;type:integer;primaryKey;autoIncrement:true;comment:ID пользователя" json:"id"`                                                             // ID пользователя
	Name             string     `gorm:"column:name;type:character varying(100);not null;comment:Имя пользователя" json:"name"`                                                              // Имя пользователя
	Email            string     `gorm:"column:email;type:character varying(100);not null;comment:Электронная почта пользователя" json:"email"`                                              // Электронная почта пользователя
	PasswordHash     string     `gorm:"column:password_hash;type:character varying(255);not null;comment:Хэш пароля пользователя" json:"password_hash"`                                     // Хэш пароля пользователя
	Phone            *string    `gorm:"column:phone;type:character varying(20);comment:Телефон пользователя" json:"phone"`                                                                  // Телефон пользователя
	Address          *string    `gorm:"column:address;type:character varying(255);comment:Адрес пользователя" json:"address"`                                                               // Адрес пользователя
	RoleID           int32      `gorm:"column:role_id;type:integer;not null;comment:ID роли пользователя" json:"role_id"`                                                                   // ID роли пользователя
	RegistrationDate *time.Time `gorm:"column:registration_date;type:timestamp without time zone;default:CURRENT_TIMESTAMP;comment:Дата регистрации пользователя" json:"registration_date"` // Дата регистрации пользователя
	LastLogin        *time.Time `gorm:"column:last_login;type:timestamp without time zone;comment:Дата последнего входа пользователя" json:"last_login"`                                    // Дата последнего входа пользователя
	AccountStatus    *bool      `gorm:"column:account_status;type:boolean;default:true;comment:Статус аккаунта (активен/заблокирован)" json:"account_status"`                               // Статус аккаунта (активен/заблокирован)
	OrderCount       *int32     `gorm:"column:order_count;type:integer;comment:Количество заказов пользователя" json:"order_count"`                                                         // Количество заказов пользователя
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}
