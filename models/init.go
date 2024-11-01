package models

func GetModels() []interface{} {
	return []interface{}{
		&UserRole{},
		&NewsletterSubscription{},
		&Manufacturer{},
		&Promotion{},
		&OrderStatus{},
		&ShoppingCart{},
		&NotificationStatus{},
		&OrderItem{},
		&CartItem{},
		&ProductChange{},
		&Payment{},
		&SystemLog{},
		&CustomerSupport{},
		&InventoryStatus{},
		&UserSession{},
		&PromotionStatus{},
		&Category{},
		&ReturnStatus{},
		&Return{},
		&Sale{},
		&Review{},
		&PaymentStatus{},
		&User{},
		&Order{},
		&Product{},
		&Notification{},
	}
}
