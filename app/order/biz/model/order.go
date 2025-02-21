package model

import (
	"gorm.io/gorm"
	"time"
)

type Order struct {
	Base
	ID              int64      `gorm:"primaryKey;autoIncrement;column:id;comment:订单ID"`
	UserID          int64      `gorm:"index;column:user_id;comment:购买用户ID"`
	TotalPrice      int64      `gorm:"column:total_price;not null;comment:订单总金额"`
	DiscountPrice   int64      `gorm:"column:discount_price;default:0;comment:优惠金额"`
	ActualPrice     int64      `gorm:"column:actual_price;not null;comment:实际支付金额"`
	OrderStatus     int8       `gorm:"column:order_status;default:0;comment:订单状态（0-待支付, 1-已支付, 2-已发货, 3-已完成, 4-已取消）"`
	PaymentStatus   int8       `gorm:"column:payment_status;default:0;comment:支付状态（0-未支付, 1-已支付, 2-支付失败, 3-退款中, 4-已退款）"`
	PaymentMethod   string     `gorm:"column:payment_method;size:20;comment:支付方式（微信、支付宝、银行卡等）"`
	ShippingAddress string     `gorm:"column:shipping_address;not null;comment:收货地址"`
	RecipientName   string     `gorm:"column:recipient_name;not null;comment:收件人姓名"`
	PhoneNumber     string     `gorm:"column:phone_number;not null;size:20;comment:收件人电话号码"`
	ShippingStatus  int8       `gorm:"column:shipping_status;default:0;comment:物流状态（0-未发货, 1-已发货, 2-已签收）"`
	PaidAt          *time.Time `gorm:"column:paid_at;comment:订单支付时间"`
	ShippedAt       *time.Time `gorm:"column:shipped_at;comment:发货时间"`
	CompletedAt     *time.Time `gorm:"column:completed_at;comment:订单完成时间"`
	CanceledAt      *time.Time `gorm:"column:canceled_at;comment:订单取消时间"`
	Remark          *string    `gorm:"column:remark;comment:订单备注"`
}

func (u Order) TableName() string {
	return "orders"
}

func GetProductIdsFromOrder(db *gorm.DB, order string) []int64 {
	return []int64{}
}
