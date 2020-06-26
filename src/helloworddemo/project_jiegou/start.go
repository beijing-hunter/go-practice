package main

import (
	"fmt"
	"project_jiegou/utils" //project_jiegou是在go.mode文件中定义的
)

func main() {

	var value1, value2 = 12.2, 1.2
	var oper byte = '+'
	var result = utils.OperFunc(value1, value2, oper)
	fmt.Printf("result=%.2f\n", result)

	var valueInt = utils.StringToFloat64("12312.1552")
	fmt.Printf("valueInt=%.2f\n", valueInt)

	count := 10
	for index := 0; index < count; index++ {
		fmt.Println("for index=", index)
	}

	for {

		if count > 100 {
			fmt.Printf("我已经%d岁了\n", count)
			break
		}

		fmt.Printf("今年%d岁\n", count)
		count++
	}

	var valueFloat64 = 12.106
	var strFloat64 = utils.Float64ToString(valueFloat64)
	fmt.Println(strFloat64)
}
