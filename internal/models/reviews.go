package models

import (
	"time"
)

const TableNameReview = "reviews"

// Review mapped from table <reviews>
type Review struct {
	ID         int32      `gorm:"column:id;type:integer;primaryKey;autoIncrement:true;comment:ID отзыва" json:"id"`                                                // ID отзыва
	UserID     int32      `gorm:"column:user_id;type:integer;not null;comment:ID пользователя, который оставил отзыв" json:"user_id"`                              // ID пользователя, который оставил отзыв
	ProductID  int32      `gorm:"column:product_id;type:integer;not null;comment:ID продукта, на который оставлен отзыв" json:"product_id"`                        // ID продукта, на который оставлен отзыв
	Rating     int32      `gorm:"column:rating;type:integer;not null;comment:Оценка отзыва (в звёздах)" json:"rating"`                                             // Оценка отзыва (в звёздах)
	ReviewText *string    `gorm:"column:review_text;type:text;comment:Текст отзыва" json:"review_text"`                                                            // Текст отзыва
	ReviewDate *time.Time `gorm:"column:review_date;type:timestamp without time zone;default:CURRENT_TIMESTAMP;comment:Дата добавления отзыва" json:"review_date"` // Дата добавления отзыва
	StatusID   int32      `gorm:"column:status_id;type:integer;not null;comment:ID статуса отзыва" json:"status_id"`                                               // ID статуса отзыва
}

// TableName Review's table name
func (*Review) TableName() string {
	return TableNameReview
}
