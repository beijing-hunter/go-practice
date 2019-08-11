package main

import "fmt"

func main(){

	intChan:=make(chan int,100)

	for i:=1;i<=50;i++{
		intChan<-i
	}

	for len(intChan)>0{
		fmt.Println(<-intChan)
	}
}
