package main

import "fmt"

func main() {

	count := 10

	for index := 0; index < count; index++ {

		fmt.Println("hello world")

		if index == 5 {
			fmt.Println("hello world", index)
		}
	}

	for { //go语言没有while,do...while的语法，注意。

		if count == 100 {
			fmt.Println("count 100 了")
			break
		} else {
			fmt.Println("go语言没有while,do...while的语法，注意。")
		}

		count++
	}

	var str = "hello,world!商河"
	for i := 0; i < len(str); i++ { //如果字符串含有中文，那么传统的遍历字符串方式，就是错误，会出现乱码。
		fmt.Printf("str %c\n", str[i])
	}

	for index, val := range str { //这个go 特有，一般使用这个遍历字符串
		fmt.Printf("index=%d,val=%c\n", index, val)
	}
}
