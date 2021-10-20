package service

import (
	"github.com/xielizyh/ctid_service/internal/model"
	"github.com/xielizyh/ctid_service/pkg/app"
)

type CountOrderRequest struct {
	Name  string `form:"name" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type OrderListRequest struct {
	Name  string `form:"name" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type CreateOrderRequest struct {
	Name      string `form:"name" binding:"required,min=3,max=100"`
	CreatedBy string `form:"created_by" binding:"required,min=3,max=100"`
	State     uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type UpdateOrderRequest struct {
	ID         uint32 `form:"id" binding:"required,gte=1"`
	Name       string `form:"name" binding:"min=3,max=100"`
	State      uint8  `form:"state" binding:"oneof=0 1"`
	ModifiedBy string `form:"modified_by" binding:"required,min=3,max=100"`
}

type DeleteOrderRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}

func (svc *Service) CountOrder(param *CountOrderRequest) (int, error) {
	return svc.dao.CountOrder(param.Name, param.State)
}

func (svc *Service) GetOrderList(param *OrderListRequest, pager *app.Pager) ([]*model.Order, error) {
	return svc.dao.GetOrderList(param.Name, param.State, pager.Page, pager.PageSize)
}

func (svc *Service) CreateOrder(param *CreateOrderRequest) error {
	return svc.dao.CreateOrder(param.Name, param.State, param.CreatedBy)
}

func (svc *Service) UpdateOrder(param *UpdateOrderRequest) error {
	return svc.dao.UpdateOrder(param.ID, param.Name, param.State, param.ModifiedBy)
}

func (svc *Service) DeleteOrder(param *DeleteOrderRequest) error {
	return svc.dao.DeleteOrder(param.ID)
}
