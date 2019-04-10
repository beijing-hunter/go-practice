package main

import "fmt"

//全局变量
var (
	pid  = 21
	pkey = "sxafewef"
)

func main() {
	var i int
	i = 10
	fmt.Println("i=", i)

	var num = 10.11
	fmt.Println("num=", num)

	num1 := 12.23
	fmt.Println("num1=", num1)

	var n1, name, n3 = 12, "sb", 12.23
	fmt.Println("n1=", n1, "name=", name, "n3=", n3)

	fmt.Println("pid=", pid, "pkey=", pkey)
}
