package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/xielizyh/ctid_service/global"
	"github.com/xielizyh/ctid_service/internal/service"
	"github.com/xielizyh/ctid_service/pkg/app"
	"github.com/xielizyh/ctid_service/pkg/convert"
	"github.com/xielizyh/ctid_service/pkg/errcode"
)

type Order struct{}

// NewOrder 新建订单
func NewOrder() Order {
	return Order{}
}

// @Summary 获取订单
// @Produce  json
// @Param id path int true "订单 ID"
// @Param state body int false "状态" Enums(0, 1) default(1)
// @Success 200 {object} model.Order "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/orders [get]
func (o Order) Get(c *gin.Context) {
	// 入参校验和参数绑定
	param := service.GetOrderRequest{
		ID: convert.StrTo(c.Param("id")).MustUInt32(),
	}
	// 创建响应
	response := app.NewResponse(c)
	// 校验
	valid, errs := app.BindAndValid(c, &param, "")
	if !valid {
		global.Logger.Errorf("app.BindAndValid error: %v", errs)
		// 回应错误
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	// 正常回应
	svc := service.New(c.Request.Context())
	Orders, err := svc.GetOrder(&param)
	if err != nil {
		global.Logger.Errorf("svc.GetOrder error: %v", err)
		response.ToErrorResponse(errcode.ErrorGetOrderListFail)
		return
	}

	response.ToResponseList(Orders, len(Orders))
}

// @Summary 获取多个订单
// @Produce  json
// @Param state body int false "状态" Enums(0, 1) default(1)
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} model.Order "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/orders [get]
func (o Order) List(c *gin.Context) {
	// 入参校验和参数绑定
	param := service.OrderListRequest{}
	// 创建响应
	response := app.NewResponse(c)
	// 校验
	valid, errs := app.BindAndValid(c, &param, "")
	if !valid {
		global.Logger.Errorf("app.BindAndValid error: %v", errs)
		// 回应错误
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	// 正常回应
	svc := service.New(c.Request.Context())
	pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	totalRows, err := svc.CountOrder(&service.CountOrderRequest{State: param.State})
	if err != nil {
		global.Logger.Errorf("svc.CountOrder error: %v", err)
		response.ToErrorResponse(errcode.ErrorCountOrderFail)
		return
	}

	Orders, err := svc.GetOrderList(&param, &pager)
	if err != nil {
		global.Logger.Errorf("svc.GetOrderList error: %v", err)
		response.ToErrorResponse(errcode.ErrorGetOrderListFail)
		return
	}
	response.ToResponseList(Orders, totalRows)
}

// @Summary 新增订单
// @Produce  json
// @Success 200 {object} model.Order "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/orders [post]
func (o Order) Create(c *gin.Context) {
	// 注意读取了后已无数据，再入场校验和参数绑定会出错
	// buf := make([]byte, 1024)
	// n, _ := c.Request.Body.Read(buf)
	// log.Println(string(buf[0:n]))
	// 入参校验和参数绑定
	param := service.CreateOrderRequest{
		// 默认创建可用订单
		State: 1,
	}
	// 创建响应
	response := app.NewResponse(c)
	// 校验
	valid, errs := app.BindAndValid(c, &param, "")
	if !valid {
		global.Logger.Errorf("app.BindAndValid error: %v", errs)
		// 回应错误
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	// 正常回应
	svc := service.New(c.Request.Context())
	// 去重
	checkParam := service.CheckOrderRequest{
		UserName:     param.UserName,
		CertNumber:   param.CertNumber,
		Phone:        param.Phone,
		RoomNumber:   param.RoomNumber,
		CheckinTime:  param.CheckinTime,
		CheckoutTime: param.CheckoutTime,
	}
	if totalRows, err := svc.CheckOrder(&checkParam); err != nil {
		global.Logger.Errorf("svc.CheckOrder error: %v", err)
		response.ToErrorResponse(errcode.ErrorCheckOrderFail)
		return
	} else {
		if totalRows != 0 {
			response.ToErrorResponse(errcode.ErrorAlreadyExistOrderFail)
			return
		}
	}

	err := svc.CreateOrder(&param)
	if err != nil {
		global.Logger.Errorf("svc.CreateOrder error: %v", err)
		response.ToErrorResponse(errcode.ErrorCreateOrderFail)
		return
	}

	response.ToResponse(gin.H{})
}

// @Summary 删除订单
// @Produce  json
// @Param id path int true "订单 ID"
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/orders/{id} [delete]
func (o Order) Delete(c *gin.Context) {
	// 入参校验和参数绑定
	param := service.DeleteOrderRequest{ID: convert.StrTo(c.Param("id")).MustUInt32()}
	// 创建响应
	response := app.NewResponse(c)
	// 校验
	valid, errs := app.BindAndValid(c, &param, "")
	if !valid {
		global.Logger.Errorf("app.BindAndValid error: %v", errs)
		// 回应错误
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	// 正常回应
	svc := service.New(c.Request.Context())
	err := svc.DeleteOrder(&param)
	if err != nil {
		global.Logger.Errorf("svc.DeleteOrder error: %v", err)
		response.ToErrorResponse(errcode.ErrorDeleteOrderFail)
		return
	}

	response.ToResponse(gin.H{})
}

// @Summary 校验订单
// @Produce  json
// @Success 200 {object} model.Order "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/orders [get]
func (o Order) Check(c *gin.Context) {
	// 入参校验和参数绑定
	param := service.CheckOrderRequest{}
	// 创建响应
	response := app.NewResponse(c)
	// 校验
	valid, errs := app.BindAndValid(c, &param, "json")
	if !valid {
		global.Logger.Errorf("app.BindAndValid error: %v", errs)
		// 回应错误
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	// 正常回应
	svc := service.New(c.Request.Context())
	totalRows, err := svc.CheckOrder(&param)
	if err != nil {
		global.Logger.Errorf("svc.CheckOrder error: %v", err)
		response.ToErrorResponse(errcode.ErrorCheckOrderFail)
		return
	}
	if totalRows > 0 {
		response.ToResponse(gin.H{})
	} else {
		response.ToErrorResponse(errcode.ErrorCheckOrderFail)
	}
}
