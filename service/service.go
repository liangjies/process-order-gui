package service

import (
	"fmt"
	"process-order/model"
	"process-order/sql"
)

// 获取订单
func GetOrderList() (result []model.OrderList, err error) {
	res, err := sql.GetOrderList()
	if err != nil {
		fmt.Println(err)
		return result, err
	}
	var drawNoList []string
	for _, v := range res {
		drawNoList = append(drawNoList, v.DrawNo)
	}
	// 查询机加工图号
	drawINProcessCold, err := sql.GeDrawINProcessCold(drawNoList)
	if err != nil {
		fmt.Println(err)
		return result, err
	}
	// 查询热加工图号
	drawINProcessHot, err := sql.GeDrawINProcessHot(drawNoList)
	if err != nil {
		fmt.Println(err)
		return result, err
	}
	// 查询PLM系统
	drawINPLM, err := sql.GeDrawINPLM(drawNoList)
	if err != nil {
		fmt.Println(err)
		return result, err
	}
	fmt.Println("drawINProcessCold", drawINProcessCold)
	for _, v := range res {
		// 系统存在则跳过
		if SliceIsExist(drawINPLM, v.DrawNo) || (SliceIsExist(drawINProcessCold, v.DrawNo) && SliceIsExist(drawINProcessHot, v.DrawNo)) {
			continue
		}
		if !SliceIsExist(drawINProcessCold, v.DrawNo) && SliceIsExist(drawINProcessHot, v.DrawNo) {
			v.Status = "机加工未完成"
			result = append(result, v)
			continue
		}
		if SliceIsExist(drawINProcessCold, v.DrawNo) && !SliceIsExist(drawINProcessHot, v.DrawNo) {
			v.Status = "热加工未完成"
			result = append(result, v)
			continue
		}
		v.Status = "未完成"
		result = append(result, v)
	}
	return result, nil
}

func SliceIsExist(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}
