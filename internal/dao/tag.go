package dao

import (
	"github.com/xielizyh/ctid_service/internal/model"
	"github.com/xielizyh/ctid_service/pkg/app"
)

func (d *Dao) CountOrder(name string, state uint8) (int, error) {
	order := model.Order{Name: name, State: state}
	return order.Count(d.engine)
}

func (d *Dao) GetOrderList(name string, state uint8, page, pageSize int) ([]*model.Order, error) {
	order := model.Order{Name: name, State: state}
	pageOffset := app.GetPageOffset(page, pageSize)
	return order.List(d.engine, pageOffset, pageSize)
}

func (d *Dao) CreateOrder(name string, state uint8, createdBy string) error {
	order := model.Order{
		Name:  name,
		State: state,
		Model: &model.Model{CreatedBy: createdBy},
	}

	return order.Create(d.engine)
}

func (d *Dao) UpdateOrder(id uint32, name string, state uint8, modifiedBy string) error {
	order := model.Order{
		Model: &model.Model{ID: id},
	}
	values := map[string]interface{}{
		"state":       state,
		"modified_by": modifiedBy,
	}
	if name != "" {
		values["name"] = name
	}
	return order.Update(d.engine, values)
}

func (d *Dao) DeleteOrder(id uint32) error {
	order := model.Order{Model: &model.Model{ID: id}}
	return order.Delete(d.engine)
}
