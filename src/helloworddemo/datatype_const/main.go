package main

import "fmt"

const (
	a = iota //一行自增+1,如果多个变量在一行，则不自增
	b
	c
	d
)

const (
	Wait          = iota //一行自增+1,如果多个变量在一行，则不自增
	Success, Fail = iota, iota
)

func main() {

	const v = 12
	fmt.Println(v)
	fmt.Println(a, b, c, d)

	fmt.Println(Wait, Success, Fail)
}
