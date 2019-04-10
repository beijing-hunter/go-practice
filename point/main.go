package main

import "fmt"

func main() {
	var i = 12
	fmt.Printf("i内存地址：%v,value=%v\n", &i, i)

	var ptr *int = &i //指针变量ptr
	fmt.Printf("ptr value=%v,ptr ==> value %v\n", ptr, *ptr)

	*ptr = 13
	fmt.Printf("i内存地址：%v,value=%v\n", &i, i)
}
