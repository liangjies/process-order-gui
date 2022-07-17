package global

import (
	"process-order/model"

	"gorm.io/gorm"
)

var (
	SYS_DB    *gorm.DB
	OrderList []model.OrderList
)
