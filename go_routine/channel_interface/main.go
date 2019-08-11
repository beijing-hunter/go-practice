package main

import "fmt"

type Student struct {
	Name string
	Age int
}

func main()  {

	var allChan chan interface{}
	allChan=make(chan interface{},10)

	allChan<-12.3
	allChan<-"刘邦"
	allChan<-Student{Name:"张飞",Age:12}
	<-allChan
	<-allChan
	stu:=<-allChan

	fmt.Println(stu.(Student).Name)//注意：需要断言
}
