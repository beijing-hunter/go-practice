package main

import "fmt"

type Student struct{
	Name string
	Age int
}

func(stu *Student) ShowInfo() string{
	str:=fmt.Sprintf("name=[%v],age=[%v]",stu.Name,stu.Age)
	return str
}

type Pupil struct{
	Student //嵌入了student匿名结构体 ，从而实现继承
}

func(p *Pupil) schollName() string{
	return "小学"
}

func main()  {
	
	//第一种使用
	p:=&Pupil{}
	p.Name="牛魔"
	p.Age=124

	fmt.Println(p.ShowInfo())

	//第二种使用
	p2:=&Pupil{}
	p2.Student.Name="牛魔2"
	p2.Student.Age=1241

	fmt.Println(p2.Student.ShowInfo())


	p3:=&Pupil{Student{"张飞",123,}}
	fmt.Println(p3.ShowInfo())
}