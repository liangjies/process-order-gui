package sql

import (
	"process-order/model"
	"time"

	"process-order/global"
)

// 获取订单列表
func GetOrderList() (res []model.OrderList, err error) {
	nowTime := time.Now()
	dDate := nowTime.AddDate(0, -1, -15)
	db := global.SYS_DB.Model(&model.OrderList{})
	err = db.Raw("select SaleOrderSQ.id as ordeID,cinvstd as drawNo,cinvname as partNo,cinvdefine6 as customerPartNo,cinvdefine4  as customer,ivouchstate as orderStatus,dDate as orderDate from [192.168.0.15].[UFDATA_111_2017].[dbo].SaleOrderSQ left join [192.168.0.15].[UFDATA_111_2017].[dbo].SaleOrderQ ON SaleOrderSQ.id=SaleOrderQ.id where (dDate >= ?) and ivouchstate='Approved' and cinvstd not like 'W-%' order by dDate", dDate).Scan(&res).Error
	return res, err
}

// 查询机加工图号
func GeDrawINProcessCold(drawNo []string) (res []string, err error) {
	db := global.SYS_DB.Model(&model.OrderList{})
	err = db.Raw("select distinct 图号 FROM [SQLSERVER].[cappdatabase].[dbo].[ProcessData] where 图号 in ?", drawNo).Scan(&res).Error
	return res, err
}

// 查询热加工图号
func GeDrawINProcessHot(drawNo []string) (res []string, err error) {
	db := global.SYS_DB.Model(&model.OrderList{})
	err = db.Raw("select distinct 图号 FROM [SQLSERVER].[cappdatabase].[dbo].[HeatPrDataA] where 图号 in ?", drawNo).Scan(&res).Error
	return res, err
}

// 查询PLM系统
func GeDrawINPLM(drawNo []string) (res []string, err error) {
	db := global.SYS_DB.Model(&model.OrderList{})
	err = db.Raw("select distinct identityno as 图号 FROM [PLM].[PLM].[dbo].[p_part] where identityno in ?", drawNo).Scan(&res).Error
	return res, err
}
