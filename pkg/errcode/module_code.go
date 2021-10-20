package errcode

var (
	ErrorGetOrderListFail = NewError(20010001, "获取订单列表失败")
	ErrorCreateOrderFail  = NewError(20010002, "创建订单失败")
	ErrorUpdateOrderFail  = NewError(20010003, "更新订单失败")
	ErrorDeleteOrderFail  = NewError(20010004, "删除订单失败")
	ErrorCountOrderFail   = NewError(20010005, "统计订单失败")
)
