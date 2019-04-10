package main

import "fmt"

func main() {

	var a, b int

	fmt.Println("请输出a b的值")
	fmt.Scanf("%d %d", &a, &b)

	if a > b {
		fmt.Println("a>b")
	} else if a < b {
		fmt.Println("a<b")
	} else {
		fmt.Println("a==b")
	}
}
