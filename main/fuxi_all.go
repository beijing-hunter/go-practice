package main

import (
	"fmt"
	"strconv"
)

var stuMap map[string]Student

func init() {
	stuMap = make(map[string]Student, 10)
}

type Student struct {
	name string
	age  int8
}

type Person struct {
	Student
	fuse string
}

type IService interface {
	save(info Student) bool
	getInfoByName(name string) Student
}

type ServiceImpl struct {
}

func (stu *ServiceImpl) save(info Student) bool {

	_, ok := stuMap[info.name]

	if ok {
		return false
	}

	stuMap[info.name] = info
	return true
}

func (stu *ServiceImpl) getInfoByName(name string) Student {
	return stuMap[name]
}

func (stu *Student) setName(name string) {
	stu.name = name
}

func main() {

	m := 14
	var n int64 = 12
	var f, a = 5.15, "hello world"

	fmt.Println(n, "f=", f)
	fmt.Println(m, "a=", a)

	f = float64(m)

	fmt.Printf("m=%d,f=%0.1f,a=%v,a datatype=%T \n", m, f, a, a)

	a = a + strconv.FormatInt(n, 10)
	fmt.Printf("m=%d,f=%0.1f,a=%v,a datatype=%T \n", m, f, a, a)

	var ages [5]int64
	ages[0] = 12

	for index, value := range ages {
		fmt.Printf("index=%d,value=%d\n", index, value)
	}

	agess := [...]float64{12, 14, 7: 14}
	agess[6] = 23

	for index, value := range agess {
		fmt.Printf("index=%d,value=%0.2f\n", index, value)
	}

	var ff *float64
	ff = &f
	*ff = f + 12.45

	fmt.Printf("ff value=%0.2f,ff address=%p\n", *ff, ff)

	var dataMap map[string]int64
	dataMap = make(map[string]int64)

	dataMap["12345"] = 12
	dataMap["1234"] = 45
	dataMap["123"] = 453
	fmt.Println(dataMap)

	delete(dataMap, "123")

	val, isFindSuccess := dataMap["123"]

	if isFindSuccess {
		fmt.Printf("key=%v,value=%d\n", "123", val)
	}

	dataP := make(map[string]int)
	fmt.Println(dataP)

	mm := 0
	mm, fff := 12, 13
	mm, fff = fff, mm

	as := [...]int{1, 2, 3}
	fmt.Printf("as datatype=%T,as value=%v\n", as, as)

	const y = 12
	const yf = 13.2
	fmt.Println(y)

	aslice := as[1:2]
	aslice[0] = 5
	aslice = append(aslice, 6)
	aslice = append(aslice, 7)

	for i := 12; i < 24; {
		aslice = append(aslice, i)
		i++
	}

	aslice[0] = 90
	fmt.Printf("aslice datatype=%T,aslice value=%v\n", aslice, aslice)
	fmt.Printf("as datatype=%T,as value=%v\n", as, as)

	var mf myFunc = getNumber
	v1, v2 := mf()
	fmt.Printf("mf datatype=%T,return=%v,return=%v\n", mf, v1, v2)

	var mff = appendFix("FD")
	fmt.Println(mff("123"))
	fmt.Println(mff("123"))
	fmt.Println(mff("123"))
	fmt.Println(mff("123"))
	fmt.Println(mff("123"))

	stu := Student{age: 1, name: "test"}
	fmt.Printf("stu datatype=%T\n", stu)

	stu2 := new(Student)
	fmt.Printf("stu2 datatype=%T\n", stu2)

	name := "test2"
	stu2.setName(name)
	fmt.Println(*stu2)

	person := new(Person)
	person.Student.age = 1
	person.Student.name = "bb"
	person.fuse = "黑色"
	fmt.Println(*person)

	serviceImpl := new(ServiceImpl)
	var iservice IService = serviceImpl
	iservice.save(stu)
	iservice.save(*stu2)

	fmt.Println(stuMap)

	stu3 := iservice.getInfoByName("test")
	stu3.name = "test4"
	fmt.Println(stu3)
	fmt.Println(stuMap)

	_, ok := iservice.(*ServiceImpl)

	if ok {
		fmt.Println("断言成功")
	}
}

func getSum(n int64, m int64) int64 {
	return n + m
}

func getNumber() (n int64, m float64) {
	return 12, 12.23
}

func appendFix(value string) func(v string) string {

	return func(vv string) string {
		value = value + vv
		return value
	}
}

type myFunc func() (int64, float64) //函数数据类型
