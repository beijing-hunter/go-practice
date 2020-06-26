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

	var stu Student //1.初始化,stu 值类型变量
	stu.Name = "智障了"
	stu.lables = append(stu.lables, "阳光")
	stu.lables = append(stu.lables, "灿烂")
	stu.lables = append(stu.lables, "日子")
	fmt.Printf("lables type=%T,value=%v,len=%v,cap=%v\n", stu.lables, stu.lables, len(stu.lables), cap(stu.lables))

	fmt.Println("stu=", stu)
	setAge(stu)
	fmt.Printf("stu=%v,stu type=%T\n", stu, stu)

	var stu2 = Student{Age: 1, Name: "字符画"} //2.初始化,stu2 值类型变量
	fmt.Printf("stu2=%v,stu2 type=%T\n", stu2, stu2)

	stu3 := new(Student) //3.初始化，stu3指针变量
	fmt.Printf("stu3=%v,stu3 type=%T\n", stu3, stu3)
}
