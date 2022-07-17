package sql

import (
	"process-order/model"

	"process-order/global"
)

// 获取订单列表
func GetOrderList() (res []model.OrderList, err error) {
	db := global.SYS_DB.Model(&model.OrderList{})
	err = db.Raw("select top 100 cinvstd as 图号,cinvname as 零件号,cinvdefine6 as 客户零件号,cinvdefine4  as 客户,ivouchstate as 状态,dDate as 订单日期 from [192.168.0.15].[UFDATA_111_2017].[dbo].SaleOrderSQ left join [192.168.0.15].[UFDATA_111_2017].[dbo].SaleOrderQ ON SaleOrderSQ.id=SaleOrderQ.id").Scan(&res).Error
	return res, err
}
