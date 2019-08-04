package main

import "fmt"

type Student struct{

	Name string 
	Age int 
}

func(s *Student) getName() string{//这个方法与Student数据类型绑定
	s.Name="甄姬"
	return s.Name
}

func(s *Student) String() string{
	str:=fmt.Sprintf("name=[%v],age=[%v]",s.Name,s.Age)
	return str
}

func main()  {
	
	stu:=Student{Name:"云中君"}
	fmt.Println("stu.getName=",stu.getName())
	fmt.Println("stu.Name=",stu.Name)

	fmt.Println(&stu)//默认调用Student的String()方法
}