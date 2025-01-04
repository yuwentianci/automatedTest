package main

import (
	"fmt"
	"myapp/secondFuture"
)

func main() {
	email := []string{
		"yuwentiangong@hotmail.com",
		"yuwenlongxiang@hotmail.com",
		"yuwenlongming@hotmail.com",
		"yuwenbaiqi@hotmail.com",
		"yuwenhanxin@hotmail.com",
		"yuwensunwu@hotmail.com",
		"yuwenzhangliang@hotmail.com",
		"yuwenzhenguan@hotmail.com",
		"yuwenzhulong@hotmail.com",
		"yuwenaolong@hotmail.com",
		"yuwenhuilong@hotmail.com",
		"yuwenshenlong@hotmail.com",
		"yuwenqilin@hotmail.com",
		"yuwenlongma@hotmail.com",
		"yuwenqiulong@hotmail.com",
		"yuwenpanlong@hotmail.com",
		"yuwenqilong@hotmail.com",
		"yuwenyinglong@hotmail.com",
		"yuwenqinglong@hotmail.com",
		"yuwenlongwang@hotmail.com",
		"yuwenhuanglong@hotmail.com",
		"yuwenjinlong@hotmail.com",
		"yuwenbixi@hotmail.com",
		"yuwenchiwen@hotmail.com",
		"yuwenpulao@hotmail.com",
		"yuwenbian@hotmail.com",
		"yuwentaotie@hotmail.com",
		"yuwenbaxia@hotmail.com",
		"yuwenyazi@hotmail.com",
		"yuwensuanni@hotmail.com",
		"yuwenjiaotu@hotmail.com"}

	err, code := secondFuture.MakePosition(1718, 1, 0, email)
	if err != nil {
		fmt.Println("报错信息:", err)
	}
	fmt.Println(code)

	//err, d := spot.Symbol()
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(d)
}
