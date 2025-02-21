package model

import (
	"fmt"
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
	Status        uint8     `json:"status"`
}

// 创建订单信息
func CreatePayment(db *gorm.DB, payment *Payment) error {
	return db.Model(Payment{}).Table("payments").Create(payment).Error
}

// 查询订单信息
func QueryPayment(db *gorm.DB, id int) (Payment, error) {
	var row Payment
	err := db.Model(&Payment{}).Where("id = ?", id).First(&row).Error
	if err != nil {
		return row, fmt.Errorf("fail to query payment %v", err)
	}
	return row, nil
}

// 更新订单信息
func UpdatePaymentStatus(db *gorm.DB, payment Payment) error {
	// 根据 ID 查找要更新的记录
	result := db.Model(&Payment{}).Where("id = ?", payment.ID).Updates(map[string]interface{}{
		"status": payment.Status,
	})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
