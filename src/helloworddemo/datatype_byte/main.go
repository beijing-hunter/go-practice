package main

import "fmt"

func main() {

	var i byte = 'a'

	fmt.Println("i=", i) //输出的ascll表码值

	fmt.Printf("i=%c\n", i) //如果希望输出对应字符，需要使用格式化输出

	var c = '国'
	cc := c + 1

	fmt.Printf("c对应的码值：%d；c=%c\n", c, c)
	fmt.Printf("cc对应的码值：%d；cc=%c\n", cc, cc)
}
