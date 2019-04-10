package main

import (
	"fmt"
	"strconv"
)

func main() {

	//基础类型 转string
	var value = 12
	var b = true
	var strValue = fmt.Sprintf("%v", value) //基础类型转换string类型
	var strB = fmt.Sprintf("%v", b)

	fmt.Printf("str type %T,value %v\n", strValue, strValue)
	fmt.Printf("strB type %T,value %v\n", strB, strB)

	var f = 12.1252
	var strF = strconv.FormatFloat(f, 'f', 2, 64) //不懂什么意思，看官方帮助文档
	fmt.Printf("strF type %T,value %v\n", strF, strF)

	var strInt = strconv.FormatInt(int64(value), 10) //不懂什么意思，看官方帮助文档
	fmt.Printf("strInt type %T,value %v\n", strInt, strInt)

	//string 转基础类型
	var str1 = "12334"
	var intStr, _ = strconv.ParseInt(str1, 10, 64)
	fmt.Printf("intStr value=%d,type=%T\n", intStr, intStr)

	var floatStr, _ = strconv.ParseFloat(str1, 64)
	fmt.Printf("floatStr value=%.2f,type=%T\n", floatStr, floatStr)
}
