package main

import "fmt"

func main() {

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
