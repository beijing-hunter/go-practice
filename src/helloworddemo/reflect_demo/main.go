package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
	Age  int8
}

func testInterface(i interface{}) {

	rType := reflect.TypeOf(i)
	fmt.Printf("type name=%v,kind=%v\n", rType.Name(), rType.Kind())

	rValue := reflect.ValueOf(i)
	iValue := rValue.Interface() //value转换为interface

	if rType.Kind() == reflect.Int64 { //是否为int64

		v, ok := iValue.(int64) //断言

		if ok {
			n := int64(12)
			sum := n + v
			fmt.Printf("sum=%v\n", sum)
		}
	}

	if rType.Kind() == reflect.Ptr { //指针类型

		fmt.Printf("rvalue kind=%v\n", rValue.Elem().Kind())

		if rValue.Elem().Kind() == reflect.Int64 { //获取指针实际指向的数据类型
			rValue.Elem().SetInt(12) //如果要改变原有的值，传递过来的类型必须是指针类型
		}
	}
}

func testInterface2(i interface{}) {

	rType := reflect.TypeOf(i)
	fmt.Printf("type name=%v,kind=%v\n", rType.Name(), rType.Kind())

	if rType.Kind() == reflect.Struct { //类型判断

		stu, ok := i.(Student)

		if ok {
			fmt.Println(stu.Name, stu.Age)
		}
	}
}

func main() {

	var num int64 = 12
	testInterface(&num)

	stu := Student{Name: "test", Age: 1}
	testInterface2(stu)
}
