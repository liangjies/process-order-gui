package service

import (
	"fmt"
	"process-order/sql"
)

// 获取订单
func GetOrderList() {
	res, err := sql.GetOrderList()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}
