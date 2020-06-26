package main

import "fmt"

func main() {

	var v interface{}
	var f float64 = 12.3
	v = f
	dd, ok := v.(float64)

	if ok {
		fmt.Println("断言成功")
		fmt.Println(dd)
	} else {
		fmt.Println("断言失败")
	}
}
