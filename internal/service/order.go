package service

import (
	"github.com/xielizyh/ctid_service/internal/model"
	"github.com/xielizyh/ctid_service/pkg/app"
)

type CountOrderRequest struct {
	State uint8 `form:"state,default=1" binding:"oneof=0 1"`
}

type GetOrderRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}

type OrderListRequest struct {
	State uint8 `json:"state" binding:"oneof=0 1"`
}

type CreateOrderRequest struct {
	UserName     string `json:"userName" binding:"required,min=2,max=100"`
	Phone        string `json:"phone" binding:"required"`
	CertType     string `json:"certType" binding:"required"`
	CertNumber   string `json:"certNumber" binding:"required"`
	RoomNumber   uint16 `json:"roomNumber" binding:"required"`
	CheckinTime  uint32 `json:"checkinTime" binding:"required"`
	CheckoutTime uint32 `json:"checkoutTime" binding:"required"`
	CreatedBy    string `json:"createdBy" binding:"min=0,max=100"`
	State        uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type UpdateOrderRequest struct {
	ID         uint32 `form:"id" binding:"required,gte=1"`
	UserName   string `form:"user_name" binding:"min=2,max=100"`
	State      uint8  `form:"state" binding:"oneof=0 1"`
	ModifiedBy string `form:"modified_by" binding:"required,min=2,max=100"`
}

type DeleteOrderRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}

func (svc *Service) CountOrder(param *CountOrderRequest) (int, error) {
	return svc.dao.CountOrder(param.State)
}

func (svc *Service) GetOrder(param *GetOrderRequest) ([]*model.Order, error) {
	return svc.dao.GetOrder(param.ID)
}

func (svc *Service) GetOrderList(param *OrderListRequest, pager *app.Pager) ([]*model.Order, error) {
	return svc.dao.GetOrderList(param.State, pager.Page, pager.PageSize)
}

func (svc *Service) CreateOrder(param *CreateOrderRequest) error {
	return svc.dao.CreateOrder(
		param.UserName,
		param.Phone,
		param.CertType,
		param.CertNumber,
		param.RoomNumber,
		param.CheckinTime,
		param.CheckoutTime,
		param.State,
	)
}

func (svc *Service) UpdateOrder(param *UpdateOrderRequest) error {
	return svc.dao.UpdateOrder(param.ID, param.UserName, param.State, param.ModifiedBy)
}

func (svc *Service) DeleteOrder(param *DeleteOrderRequest) error {
	return svc.dao.DeleteOrder(param.ID)
}
