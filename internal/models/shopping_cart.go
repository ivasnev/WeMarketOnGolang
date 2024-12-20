package models

import (
	"time"
)

const TableNameShoppingCart = "shopping_cart"

// ShoppingCart mapped from table <shopping_cart>
type ShoppingCart struct {
	ID          int32      `gorm:"column:id;type:integer;primaryKey;autoIncrement:true;comment:ID корзины" json:"id"`                                                // ID корзины
	UserID      int32      `gorm:"column:user_id;type:integer;not null;comment:ID пользователя, которому принадлежит корзина" json:"user_id"`                        // ID пользователя, которому принадлежит корзина
	Status      *bool      `gorm:"column:status;type:boolean;default:true;comment:Статус корзины (активна/неактивна)" json:"status"`                                 // Статус корзины (активна/неактивна)
	CreatedDate *time.Time `gorm:"column:created_date;type:timestamp without time zone;default:CURRENT_TIMESTAMP;comment:Дата создания корзины" json:"created_date"` // Дата создания корзины
	PromoCode   *string    `gorm:"column:promo_code;type:character varying(50);comment:Промокод, если применён" json:"promo_code"`                                   // Промокод, если применён
	TotalAmount *float64   `gorm:"column:total_amount;type:numeric(10,2);comment:Итоговая сумма корзины" json:"total_amount"`                                        // Итоговая сумма корзины
}

// TableName ShoppingCart's table name
func (*ShoppingCart) TableName() string {
	return TableNameShoppingCart
}
