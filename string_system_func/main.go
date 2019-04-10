package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	var strValue = "hello world北京"
	fmt.Println("strValue 长度=", len(strValue)) //len是按照字节返回长度的

	for index, val := range strValue {
		fmt.Printf("index=%d,value=%c\n", index, val)
	}

	n, error := strconv.Atoi("11h1") //字符串转化成整型

	if error != nil {
		fmt.Println("转化错误：", error)
	} else {
		fmt.Println("转化成果：", n)
	}

	var bytes = []byte("hello go go")
	strValue = string(bytes) //将byte数组转换成string
	fmt.Println(strValue)

	isExist := strings.Contains(strValue, "go") //区分大小写，是否包含某个字符串
	fmt.Println(isExist)

	count := strings.Count(strValue, "Go") //区分大小写，包含几个字符串
	fmt.Println(count)

	isSuccess := strings.EqualFold("ab", "Ab") //不区分大小写，判断两个字符串是否相等
	fmt.Println(isSuccess)

	index := strings.Index(strValue, "go") //区分大小写，子串在哪个位置
	fmt.Println(index)

	strValue = strings.ReplaceAll(strValue, "go", "北京")
	fmt.Println(strValue)

	strValue = "  你好，北京!"
	fmt.Println(strings.TrimSpace(strValue))  //去掉左右两边的空格
	fmt.Println(strings.Trim(strValue, "!"))  //去掉左右两边的!
	fmt.Println(strings.Trim(strValue, " !")) //去掉左右两边的!和空格
}
