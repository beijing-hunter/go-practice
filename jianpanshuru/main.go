package main

import (
	"flag"
	"fmt"
)

var (
	handlerfuc = ""
)

func init() { //命令行传参接收：如：go run main.go -handlerfuc func
	flag.StringVar(&handlerfuc, "handlerfuc", "defaultfunc", "需要执行那个处理函数")
}

func main() {

	flag.Parse() //暂停获取参数

	fmt.Println("handlerfuc=", handlerfuc)
	var name string
	var age byte

	fmt.Println("请输入姓名")
	fmt.Scanln(&name) //换行输入

	fmt.Println("请输入年龄")
	fmt.Scanln(&age) //换行输入
	fmt.Printf("用户姓名：%v,年龄：%v\n", name, age)

	fmt.Println("请输入姓名 年龄,字段使用空格隔开")
	fmt.Scanf("%s %d", &name, &age) //格式化输入
	fmt.Printf("用户姓名：%v,年龄：%v\n", name, age)
}
