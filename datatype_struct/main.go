package main

import "fmt"

type Student struct {
	Name   string
	Age    int
	lables []string
}

func setAge(stu Student) {
	stu.Age = 100
}

func main() {

	var stu Student
	stu.Name = "智障了"
	stu.lables = append(stu.lables, "阳光")
	stu.lables = append(stu.lables, "灿烂")
	stu.lables = append(stu.lables, "日子")
	fmt.Printf("lables type=%T,value=%v,len=%v,cap=%v\n", stu.lables, stu.lables, len(stu.lables), cap(stu.lables))

	fmt.Println("stu=", stu)
	setAge(stu) //结构体是值类型
	fmt.Println("stu=", stu)
}
