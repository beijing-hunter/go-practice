package main

import "fmt"

func main() {

	//定义channel
	var strChan chan string
	strChan = make(chan string, 4)

	//channel队列写入数据
	strChan <- "马超"
	strChan <- "赵云"
	fmt.Printf("strChan len=%v,cap=%v\n", len(strChan), cap(strChan))

	//channel队列取出数据
	strName := <-strChan
	fmt.Println(strName)

	var intChan chan int64
	intChan = make(chan int64, 5)

	intChan <- 23
	intChan <- 24
	value := <-intChan
	fmt.Println(value)

}
