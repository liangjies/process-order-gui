package model

type OrderList struct {
	OrderID        string `gorm:"column:order_id"`
	DrawNo         string `gorm:"column:drawNo"`
	PartNo         string `gorm:"column:partNo"`
	CustomerPartNo string `gorm:"column:customerPartNo"`
	Customer       string `gorm:"column:customer"`
	OrderDate      string `gorm:"column:orderDate"`
}
