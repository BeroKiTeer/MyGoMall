package model

import (
	"gorm.io/gorm"
	"time"
)

type Payment struct {
	gorm.Model
	UserID        uint32    `json:"user_id"`
	OrderID       string    `json:"order_id"`
	TransactionID string    `json:"transaction_id"`
	Amount        int64     `json:"amount"`
	PayAt         time.Time `json:"pay_at"`
}

func CreatePayment(db *gorm.DB, payment *Payment) error {
	return db.Model(Payment{}).Table("payments").Create(payment).Error
}
