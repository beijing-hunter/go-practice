package main

import "time"

type Student struct {
	Name string
	Age  int8
}

func main() { //GODEBUG=gctrace=1 go run main.go

	var stus []Student

	for true {

		stu := Student{Name: "gc_test", Age: 12}
		stus = append(stus, stu)

		if len(stus) > 100000 {
			time.Sleep(10)
			//break
		}

	}

	//fmt.Println(stus)
}
