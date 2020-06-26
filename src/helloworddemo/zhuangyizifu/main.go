package main

import "fmt"

func main() {
	//制表符
	fmt.Println("gao\tya\tlei")

	// \r 回车，从当前行的最前面开始输出，覆盖掉以前内容
	fmt.Println("    张飞\r阉人")

	//换行
	fmt.Println("张飞\n刘备")
}
