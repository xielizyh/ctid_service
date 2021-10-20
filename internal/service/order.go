package service

import (
	"github.com/xielizyh/ctid_service/internal/model"
	"github.com/xielizyh/ctid_service/pkg/app"
)

type CountOrderRequest struct {
	UserName string `form:"user_name" binding:"max=100"`
	State    uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type OrderListRequest struct {
	UserName string `form:"user_name" binding:"max=100"`
	State    uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type CreateOrderRequest struct {
	UserName     string `form:"user_name" binding:"required,min=2,max=100"`
	Phone        uint64 `form:"phone" binding:"required"`
	CertType     string `form:"cert_type" binding:"required"`
	CertNumber   uint64 `form:"cert_number" binding:"required"`
	RoomNumber   uint16 `form:"room_number" binding:"required"`
	CheckinTime  uint32 `form:"checkin_time" binding:"required"`
	CheckoutTime uint32 `form:"checkout_time" binding:"required"`
	CreatedBy    string `form:"created_by" binding:"required,min=2,max=100"`
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
	return svc.dao.CountOrder(param.UserName, param.State)
}

func (svc *Service) GetOrderList(param *OrderListRequest, pager *app.Pager) ([]*model.Order, error) {
	return svc.dao.GetOrderList(param.UserName, param.State, pager.Page, pager.PageSize)
}

func (svc *Service) CreateOrder(param *CreateOrderRequest) error {
	return svc.dao.CreateOrder(param.UserName, param.State, param.CreatedBy)
}

func (svc *Service) UpdateOrder(param *UpdateOrderRequest) error {
	return svc.dao.UpdateOrder(param.ID, param.UserName, param.State, param.ModifiedBy)
}

func (svc *Service) DeleteOrder(param *DeleteOrderRequest) error {
	return svc.dao.DeleteOrder(param.ID)
}
