package main

import (
	"fmt"
	"time"
)

func show(){

	for i:=0;i<11;i++{
		fmt.Println("hello.show",i)
		time.Sleep(time.Second*2)//每隔2s输出
	}
}

func main(){

	go show() //开启一个协程

	for i:=0;i<11;i++{
		fmt.Println("hello.main",i)
		time.Sleep(time.Second)
	}
}
