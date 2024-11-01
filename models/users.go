package models

import (
	"time"
)

const TableNameUser = "users"

// User mapped from table <users>
type User struct {
	ID               int32     `gorm:"column:id;primaryKey;autoIncrement:true;comment:ID пользователя" json:"id"`                                         // ID пользователя
	Name             string    `gorm:"column:name;not null;comment:Имя пользователя" json:"name"`                                                         // Имя пользователя
	Email            string    `gorm:"column:email;not null;comment:Электронная почта пользователя" json:"email"`                                         // Электронная почта пользователя
	PasswordHash     string    `gorm:"column:password_hash;not null;comment:Хэш пароля пользователя" json:"password_hash"`                                // Хэш пароля пользователя
	Phone            string    `gorm:"column:phone;comment:Телефон пользователя" json:"phone"`                                                            // Телефон пользователя
	Address          string    `gorm:"column:address;comment:Адрес пользователя" json:"address"`                                                          // Адрес пользователя
	RoleID           int32     `gorm:"column:role_id;not null;comment:ID роли пользователя" json:"role_id"`                                               // ID роли пользователя
	RegistrationDate time.Time `gorm:"column:registration_date;default:CURRENT_TIMESTAMP;comment:Дата регистрации пользователя" json:"registration_date"` // Дата регистрации пользователя
	LastLogin        time.Time `gorm:"column:last_login;comment:Дата последнего входа пользователя" json:"last_login"`                                    // Дата последнего входа пользователя
	AccountStatus    bool      `gorm:"column:account_status;default:true;comment:Статус аккаунта (активен/заблокирован)" json:"account_status"`           // Статус аккаунта (активен/заблокирован)
	OrderCount       int32     `gorm:"column:order_count;comment:Количество заказов пользователя" json:"order_count"`                                     // Количество заказов пользователя
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}
