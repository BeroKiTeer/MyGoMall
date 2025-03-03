package model

import (
	"fmt"
	"gorm.io/gorm"
	"order/biz/dal/mysql"
	"time"
)

type Order struct {
	Base
	UserID         int64      `gorm:"index;column:user_id;comment:购买用户ID"`
	TotalPrice     int64      `gorm:"column:total_price;not null;comment:订单总金额"`
	DiscountPrice  int64      `gorm:"column:discount_price;default:0;comment:优惠金额"`
	ActualPrice    int64      `gorm:"column:actual_price;not null;comment:实际支付金额"`
	OrderStatus    int8       `gorm:"column:order_status;default:0;comment:订单状态（0-待支付, 1-已支付, 2-已发货, 3-已完成, 4-已取消）"`
	PaymentStatus  int8       `gorm:"column:payment_status;default:0;comment:支付状态（0-未支付, 1-已支付, 2-支付失败, 3-退款中, 4-已退款）"`
	PaymentMethod  string     `gorm:"column:payment_method;size:20;comment:支付方式（微信、支付宝、银行卡等）"`
	RecipientName  string     `gorm:"column:recipient_name;not null;comment:收件人姓名"`
	PhoneNumber    string     `gorm:"column:phone_number;not null;size:20;comment:收件人电话号码"`
	ShippingStatus int8       `gorm:"column:shipping_status;default:0;comment:物流状态（0-未发货, 1-已发货, 2-已签收）"`
	PaidAt         *time.Time `gorm:"column:paid_at;comment:订单支付时间"`
	ShippedAt      *time.Time `gorm:"column:shipped_at;comment:发货时间"`
	CompletedAt    *time.Time `gorm:"column:completed_at;comment:订单完成时间"`
	CanceledAt     *time.Time `gorm:"column:canceled_at;comment:订单取消时间"`
	Remark         *string    `gorm:"column:remark;comment:订单备注"`
}

type OrderItem struct {
	ID        int64  `gorm:"column:id;primary_key;auto_increment;comment:订单项ID"`
	OrderId   string `gorm:"column:order_id; not null; comment:关联的订单ID"`
	ProductId int64  `gorm:"column:product_id; not null; comment:关联的商品ID"`
	Quantity  int64  `gorm:"column:quantity; not null; comment:商品数量"`
	Price     int64  `gorm:"column:price; not null; comment:商品单价"`
}

func (u Order) TableName() string {
	return "orders"
}

func GetProductIdsFromOrder(db *gorm.DB, order string) (productIds []int64, err error) {
	if err = db.Model(&OrderItem{}).Where("order_id=?", order).Pluck("product_id", &productIds).Error; err != nil {
		return nil, err
	}
	return productIds, nil
}

func UpdateOrderStatus(db *gorm.DB, orderID string, status string) {

}

func GetOrdersByUserID(db *gorm.DB, UserID int64) ([]Order, error) {
	// 检查数据库连接是否有效
	if db == nil {
		return nil, fmt.Errorf("database connection is nil")
	}

	var orders []Order
	// 查询与用户ID相关联的所有订单
	err := db.Model(&Order{}).Where("user_id=?", UserID).Find(&orders).Error
	if err != nil {
		return nil, fmt.Errorf("failed to find orders for user: %w", err)
	}
	return orders, nil
}

func GetOrderItemByOrderID(db *gorm.DB, OrderID string) ([]OrderItem, error) {
	// 检查数据库连接是否有效
	if db == nil {
		return nil, fmt.Errorf("database connection is nil")
	}
	var orderitems []OrderItem
	// 查询与用户ID相关联的所有订单
	err := db.Model(&OrderItem{}).Where("order_id=?", OrderID).Find(&orderitems).Error
	if err != nil {
		return nil, fmt.Errorf("failed to find orderitems: %w", err)
	}
	return orderitems, nil
}

func GetOrder(db *gorm.DB, ID string) (Order, error) {
	var row Order
	err := db.Model(&Order{}).Where("ID=?", ID).Find(&row).Error
	if err != nil {
		return row, fmt.Errorf("failed to find order: %w", err)
	}
	return row, nil
}

func CreateOrder(db *gorm.DB, order *Order) {
	mysql.DB.Table("orders").Create(order)
}

func CreateOrderItem(db *gorm.DB, orderItem *OrderItem) error {
	return mysql.DB.Table("order_item").Create(orderItem).Error
}

func UpdateOrder(db *gorm.DB, order *Order) error {
	return db.Exec(`update order 
					   set recipient_name=?,
					       phone_number=?,
					       order_status=?,
					       payment_status=?,
					       payment_method=?
					   where id=?`,
		order.RecipientName,
		order.PhoneNumber,
		order.OrderStatus,
		order.PaymentStatus,
		order.PaymentMethod,
		order.ID,
	).Error
}

func CancelOrder(db *gorm.DB, status int8, orderID string) error {
	return db.Exec(`update order set order_status=? where id=?`, status, orderID).Error
}

func SelectOrderItemsById(db *gorm.DB, orderID string) (orderItems []*OrderItem, err error) {
	err = db.Model(&OrderItem{}).
		Where("order_id=?", orderID).
		Find(&orderItems).Error
	return orderItems, err
}
