package main

import (
	"fmt"
	"time"
)

//生产与消费

var(
	intChan chan  int
	workOKChan chan int//工作完成channel
)

func writerData()  {

	for i:=1;i<=50;i++{
		intChan<-i
		fmt.Println("writerData=",i)
	}

	intChan<-99//告诉消费者,已经生产结束
	workOKChan<-99//生产者工作完成
}

func readData(workNum int){

	for{
		if len(intChan)>0{
			data:=<-intChan

			if data==99{
				workOKChan<-99//消费者工作完成
				break
			}

			fmt.Println("readData",workNum,"=",data)
			time.Sleep(time.Second*2)
		}
	}
}

func main()  {

	intChan=make(chan int,30)
	workOKChan=make(chan int,20)

	go writerData()
	go readData(1)
	go readData(2)
	go readData(3)

	for{
		if len(workOKChan)>=2{
			fmt.Println("工作全部完成")
			break
		}
	}
}
