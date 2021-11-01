package model

import (
	"github.com/jinzhu/gorm"
)

type Order struct {
	*Model
	UserName     string `json:"user_name"`
	Phone        string `json:"phone"`
	CertType     string `json:"cert_type"`
	CertNumber   string `json:"cert_number"`
	RoomNumber   uint16 `json:"room_number"`
	CheckinTime  uint32 `json:"checkin_time"`
	CheckoutTime uint32 `json:"checkout_time"`
	State        uint8  `json:"state"`
}

// TableName 重写TableName指定其对应返回的表名
func (o Order) TableName() string {
	return "ctid_order"
}

// Count 查找订单数量
func (o Order) Count(db *gorm.DB) (int, error) {
	var count int
	// 使用state过滤
	db = db.Where("state = ?", o.State)
	// 统计可使用的订单
	if err := db.Model(&o).Where("is_del = ?", 0).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

// Get 查询订单
func (o Order) Get(db *gorm.DB) ([]*Order, error) {
	var orders []*Order
	var err error
	// 查找可使用标签的所有记录
	if err = db.Where("id = ? AND is_del = ?", o.Model.ID, 0).Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

// List 查询订单列表
func (o Order) List(db *gorm.DB, pageOffset, pageSize int) ([]*Order, error) {
	var orders []*Order
	var err error
	if pageOffset >= 0 && pageSize > 0 {
		// 偏移并限制检索的记录数
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	// fmt.Println("order.list", o.UserName, o.CertNumber, o.Phone)
	if o.UserName != "" {
		// 使用name过滤
		db = db.Where("user_name = ?", o.UserName)
	}
	if o.CertNumber != "" {
		// 使用name过滤
		db = db.Where("cert_number = ?", o.CertNumber)
	}
	if o.Phone != "" {
		// 使用name过滤
		db = db.Where("phone = ?", o.Phone)
	}
	// 查找可使用标签的所有记录
	if err = db.Where("is_del = ?", 0).Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

// Create 创建订单
func (o Order) Create(db *gorm.DB) error {
	return db.Create(&o).Error
}

// Update 更新订单
func (o Order) Update(db *gorm.DB, values interface{}) error {
	// 使用字典形式可以进行更新零值
	if err := db.Model(o).Where("id = ? AND is_del = ?", o.ID, 0).Updates(values).Error; err != nil {
		return err
	}
	return nil
}

// Delete 删除订单
func (o Order) Delete(db *gorm.DB) error {
	return db.Where("id = ? AND is_del = ?", o.Model.ID, 0).Delete(&o).Error
}

// Check 校验订单
func (o Order) Check(db *gorm.DB) (int, error) {
	var count int
	if o.UserName != "" {
		// 使用user_name过滤
		db = db.Where("user_name = ?", o.UserName)
	}
	if o.CertNumber != "" {
		// 使用cert_number过滤
		db = db.Where("cert_number = ?", o.CertNumber)
	}
	if o.Phone != "" {
		// 使用phone过滤
		db = db.Where("phone = ?", o.Phone)
	}
	if o.RoomNumber != 0 {
		// 使用RoomNumber过滤
		db = db.Where("room_number = ?", o.RoomNumber)
	}
	if o.CheckinTime != 0 {
		// 使用RoomNumber过滤
		db = db.Where("checkin_time = ?", o.CheckinTime)
	}
	if o.CheckoutTime != 0 {
		// 使用RoomNumber过滤
		db = db.Where("checkout_time = ?", o.CheckoutTime)
	}
	// fmt.Println("order.Check", o.UserName, o.CertNumber, o.Phone)
	// 统计可使用的订单
	if err := db.Model(&o).Where("is_del = ?", 0).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
