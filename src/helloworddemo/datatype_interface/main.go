package main

import "fmt"

type IUsb interface{
	Start()
	Stop()
}

type Phone struct{

}

func(phone *Phone)Start(){
	fmt.Println("手机开始工作")
}

func(phone *Phone)Stop(){
	fmt.Println("手机停止工作")
}

type Person struct{

}

func(phone *Person)Start(){
	fmt.Println("人类开始工作")
}

func(phone *Person)Stop(){
	fmt.Println("人类停止工作")
}

func main()  {
	
	//。只要一个结构体，含有了接口类型中的所有方法，那么这个结构体就实现了这个接口
	var usb IUsb=&Phone{}
	usb.Start()
	usb.Stop()

	var pUsb IUsb=&Person{}
	pUsb.Start()
	pUsb.Stop()
}