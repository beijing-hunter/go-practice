package main

import (
	"fmt"
	"reflect"
)

type Parent struct {
	Face string `json:"face"`
}

type Student struct {
	Parent
	Name string `json:"name"`
	Age  int8   `json:"age"`
	Sex  int8
}

func (p Parent) GetFace() string {
	return p.Face
}

func (stu Student) GetName() string {
	return stu.Name
}

func (stu Student) SetAge(age int8) {

	stu.Age = age
}

func testStruct(i interface{}) {

	rType := reflect.TypeOf(i)
	rValue := reflect.ValueOf(i)

	if rType.Kind() == reflect.Ptr {

		if rType.Elem().Kind() == reflect.Struct { //指针实际指向的数据类型
			rType = rType.Elem()
			rValue = rValue.Elem()
		} else {
			return
		}
	} else {
		if rType.Kind() != reflect.Struct {
			return
		}
	}

	fieldNum := rValue.NumField() //只返回公共字段，不包括继承的字段
	fmt.Printf("field num=%v\n", fieldNum)

	for i := 0; i < fieldNum; i++ {

		field := rType.Field(i)
		tagValue := field.Tag.Get("json") //只能使用type获取字段标签值
		fieldValue := rValue.Field(i)

		if tagValue != "" {
			fmt.Printf("field name=%v,tag name=%v,field value=%v\n", field.Name, tagValue, fieldValue)
		}
	}

	methodNum := rValue.NumMethod() //只返回公共方法个数,包括继承的方法
	fmt.Printf("method num=%v\n", methodNum)

	methodValue := rValue.MethodByName("GetName").Call(nil) //rValue.Method(0).Call(nil)
	fmt.Printf("method name=%v,return value=%v\n", rType.Method(1).Name, methodValue)

	methodType, _ := rType.MethodByName("SetAge")

	var param []reflect.Value
	param = append(param, reflect.ValueOf(int8(44)))
	methodValue = rValue.MethodByName("SetAge").Call(param)
	fmt.Printf("method name=%v,return value=%v,param num=%v\n", rType.Method(2).Name, methodValue, methodType.Type.NumIn())
}

func main() {

	stu := &Student{
		Name: "狼来了",
		Age:  12,
	}

	stu.Face = "red"
	stu.GetFace()
	testStruct(stu)
	fmt.Println(stu)
}
