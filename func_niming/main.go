package main

import "fmt"

func main() {

	var result = getSub(13, 16, func(val1 int, val2 int) int {
		return val1 - val2 //匿名函数
	})

	fmt.Println("result=", result)

	var n = 10
	var funcName = func(val int) int {
		n = val + n
		return n
	}

	fmt.Println(funcName(1))
	fmt.Println(funcName(1))
	fmt.Println(funcName(1))
	fmt.Println(funcName(1))
	fmt.Println(funcName(1))
}

func getSub(val1 int, val2 int, funcName func(val1 int, val2 int) int) int {

	if val1 >= val2 {
		return funcName(val1, val2)
	} else {
		var t = val1
		val1 = val2
		val2 = t
		return funcName(val1, val2)
	}
}
