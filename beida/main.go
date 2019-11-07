package main

import (
	"beida/service"
	"fmt"
)

func main() {

	stu := service.Student{Name: "tset", Password: "1234"}
	result := service.AddInfo(stu)
	fmt.Println(result)
}
