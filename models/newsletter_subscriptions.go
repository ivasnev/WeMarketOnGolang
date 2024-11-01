package models

import (
	"time"
)

const TableNameNewsletterSubscription = "newsletter_subscriptions"

// NewsletterSubscription mapped from table <newsletter_subscriptions>
type NewsletterSubscription struct {
	ID               int32     `gorm:"column:id;primaryKey;autoIncrement:true;comment:ID подписки" json:"id"`                             // ID подписки
	Email            string    `gorm:"column:email;not null;comment:Электронная почта" json:"email"`                                      // Электронная почта
	SubscriptionDate time.Time `gorm:"column:subscription_date;default:CURRENT_TIMESTAMP;comment:Дата подписки" json:"subscription_date"` // Дата подписки
	Status           bool      `gorm:"column:status;default:true;comment:Статус подписки (активна/отписался)" json:"status"`              // Статус подписки (активна/отписался)
}

// TableName NewsletterSubscription's table name
func (*NewsletterSubscription) TableName() string {
	return TableNameNewsletterSubscription
}
