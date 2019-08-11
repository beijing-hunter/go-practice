package main

import "fmt"

func main()  {
	
	
	intChan:=make(chan int,10)
	
	for i:=0;i<=5;i++{
		intChan<-i
	}

	intChan2:=make(chan int,10)

	for i:=10;i<=17;i++{
		intChan2<-i
	}

	for{
		select {
		case v:=<-intChan://如果intChan一直没有关闭，不会一直堵塞而deadloack,会自动到下一个case匹配
			fmt.Println("intChan value=",v)
		case v:=<-intChan2:
			fmt.Println("intChan2 value=",v)
		default:
			fmt.Println("程序没有读取到数据")
			return
		}
	}
}
