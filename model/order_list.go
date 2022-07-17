package model

import "time"

type OrderList struct {
	OrderID        int       `gorm:"column:ordeID"`
	DrawNo         string    `gorm:"column:drawNo"`
	PartNo         string    `gorm:"column:partNo"`
	CustomerPartNo string    `gorm:"column:customerPartNo"`
	Customer       string    `gorm:"column:customer"`
	OrderDate      time.Time `gorm:"column:orderDate"`
	OrderStatus    string    `gorm:"column:orderStatus"`
	Status         string    `gorm:"column:status"`
}
