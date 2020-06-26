package main

import "fmt"

func main() {

	str := "hello beijing"
	var sliceByte = []byte(str) //string转成byte切片
	//sliceByte[0] = '北'          //byte中不能存放中文字符，应为1个汉字占3个byte(utf-8)
	fmt.Println("sliceByte=", sliceByte)

	var sliceRune = []rune(str)
	sliceRune[0] = '北' //rune
	str = string(sliceRune)
	fmt.Println("str=", str)
}
