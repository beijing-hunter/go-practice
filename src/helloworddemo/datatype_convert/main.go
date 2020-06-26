package main

import "fmt"

func main() {

	var i = 12
	var f = float32(i) //数据类型转换均是显示转换
	var i2 = int8(i)
	fmt.Printf("f=%f,i2=%v\n", f, i2)

	var n1 = int64(f) + 12 //数据类型
	fmt.Printf("n1=%v,n1 dataType=%T\n", n1, n1)
}
