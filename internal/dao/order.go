package dao

import (
	"github.com/xielizyh/ctid_service/internal/model"
	"github.com/xielizyh/ctid_service/pkg/app"
)

func (d *Dao) CountOrder(state uint8) (int, error) {
	order := model.Order{State: state}
	return order.Count(d.engine)
}

func (d *Dao) GetOrder(id uint32) ([]*model.Order, error) {
	order := model.Order{Model: &model.Model{ID: id}}
	return order.Get(d.engine)
}

func (d *Dao) GetOrderList(state uint8, page, pageSize int) ([]*model.Order, error) {
	order := model.Order{
		State: state,
	}
	pageOffset := app.GetPageOffset(page, pageSize)
	return order.List(d.engine, pageOffset, pageSize)
}

func (d *Dao) CreateOrder(userName, phone, certType, certNumber string, roomNumber uint16, checkinTime, checkoutTime uint32, state uint8) error {
	order := model.Order{
		UserName:     userName,
		Phone:        phone,
		CertType:     certType,
		CertNumber:   certNumber,
		RoomNumber:   roomNumber,
		CheckinTime:  checkinTime,
		CheckoutTime: checkoutTime,
		State:        state,
		Model:        &model.Model{},
	}

	return order.Create(d.engine)
}

func (d *Dao) UpdateOrder(id uint32, userName string, state uint8, modifiedBy string) error {
	order := model.Order{
		Model: &model.Model{ID: id},
	}
	values := map[string]interface{}{
		"state":       state,
		"modified_by": modifiedBy,
	}
	if userName != "" {
		values["user_name"] = userName
	}
	return order.Update(d.engine, values)
}

func (d *Dao) DeleteOrder(id uint32) error {
	order := model.Order{Model: &model.Model{ID: id}}
	return order.Delete(d.engine)
}

func (d *Dao) CheckOrder(userName string, certNumber string, phone string) (int, error) {
	order := model.Order{
		UserName:   userName,
		CertNumber: certNumber,
		Phone:      phone,
		Model:      &model.Model{},
	}
	return order.Check(d.engine)
}
