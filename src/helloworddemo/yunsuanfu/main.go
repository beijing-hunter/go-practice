package main

import "fmt"

func main() {

	fmt.Println(12 / 5)

	fmt.Println(12 / 5.0)

	var i = 10
	i++
	fmt.Println(i)

	//i=i++错误的使用方式。 注意：在golang中，++和--只能独立使用，不能参数赋值等常规操作

	var a = 12
	var b = 14

	fmt.Println("a == b", a == b)
	fmt.Println("a >= b", a >= b)
	fmt.Println("a <= b", a <= b)
	fmt.Println("a != b", a != b)

	if a > 10 && a < 13 {
		fmt.Println("ok1")
	}

	if b > 15 || b > a {
		fmt.Println("ok2")
	}

	if (b - a) > 0 {
		fmt.Println("ok3")
	}
}
