package constant

const (
	// WaitingForPay 待支付
	WaitingForPay = iota
	// Paid 已支付
	Paid
	// delivered 已发货
	delivered
	// Finished 已完成
	Finished
	// Canceled 已取消
	Canceled
)
