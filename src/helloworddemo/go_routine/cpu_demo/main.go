package main

import (
	"fmt"
	"runtime"
)

func main(){

	cpuNum:=runtime.NumCPU();
	fmt.Println("cpu num=",cpuNum)

	runtime.GOMAXPROCS(cpuNum-1)
	fmt.Println("设置运行go线程个数成功")
}
