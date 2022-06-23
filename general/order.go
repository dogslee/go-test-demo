package general

// repo
type Order interface {
	Count(orderId int) int
}

type OrderBiz struct {
	Data Order
}

func (biz *OrderBiz) GetOrderCount(orderId int) int {
	return biz.Data.Count(orderId)
}
