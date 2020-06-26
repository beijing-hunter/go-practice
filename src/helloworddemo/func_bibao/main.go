package main

import "fmt"

//闭包：函数的返回类型为函数，
func appendStr(pifx string) func(value string) string {

	return func(value string) string {
		pifx = pifx + value
		return pifx
	}
}

func main() {

	var funcName = appendStr("hello")
	fmt.Println(funcName(" world"))
	fmt.Println(funcName(" world2"))
	fmt.Println(funcName(" world3"))
}
