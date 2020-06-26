package main

import (
	"errors"
	"fmt"
	"strings"
)

func errorCatch() {
	error := recover() //异常接受内置函数
	if error != nil {
		fmt.Println(error)
	}
}

func test() {
	defer errorCatch()
	n1 := 100
	n2 := 0
	value := n1 / n2
	fmt.Println("value=", value)
}

func test2(fileName string) (err error) {

	if strings.EqualFold(fileName, "init.config") {
		fmt.Println("加载文件。。。。")
		return nil
	} else {
		return errors.New("文件名称不存在") //自定义异常
	}
}

func main() {
	test()

	error := test2("init.config2")
	if error != nil {
		panic(error) //异常抛出。
	}

	fmt.Println("main继续执行")
}
