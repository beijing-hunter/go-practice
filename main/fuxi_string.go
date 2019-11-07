package main

import (
	"fmt"
	"strconv"
)

func main1() {

	str := "hello" + "world" +
		"go"

	fmt.Printf("str=%v\n", str)

	var i int64 = 30
	f := float64(i)
	fmt.Printf("i=%v,f=%.2f\n", i, f)

	var si string = "101"
	ii, _ := strconv.ParseInt(si, 10, 64)
	fmt.Printf("ii=%v,dateType=%T\n", ii, ii)

	var ipoint *int64
	ipoint = &i
	fmt.Println(*ipoint)
	*ipoint = 45
	fmt.Println(i)

	for j := 0; j < 10; j++ {

		if (j % 2) == 0 {
			fmt.Printf("j=%d,偶数\n", j)
		}
	}

	var j int64
	fmt.Scanf("%v", &j)

	switch j {
	case 12:
		fmt.Println(12)
	case 13:
		fmt.Println(13)
	default:
		fmt.Println("default 0")
	}

	var arr [3]int64
	arr[0] = 12

	for index, value := range arr {
		fmt.Printf("index=%v,value=%v\n", index, value)
	}

	arr2 := [...]int64{2, 13, 22, 45}
	fmt.Println(arr2)

	arr3 := [...]int64{2: 12, 23: 99}
	fmt.Println(arr3)

	var maps map[string]string
	maps = make(map[string]string, 10)
	maps["key1"] = "value1"
	maps["key2"] = "value2"

	value, isReadSuccess := maps["key3"]

	if isReadSuccess {
		fmt.Printf("value=%v\n", value)
	} else {
		fmt.Println("key3 不存在")
	}

	delete(maps, "key1")
	fmt.Println(maps)

	mslipce := arr2[:]

	mslipce = append(mslipce, 10)
	mslipce[0] = 3
	fmt.Println(mslipce)
	fmt.Println(arr2)

	slipce2 := make([]float64, 4, 10)
	slipce2[0] = 12
	slipce2 = append(slipce2, 34.9)
	fmt.Println(slipce2)

	var func2 = func(a int64, b int64) int64 {
		return a + b
	}

	var sum = func2(12, 14)
	fmt.Printf("sum=%d\n", sum)

	var fd = stringBuild("P")
	fmt.Println(fd("hello2"))
	fmt.Println(fd("hello3"))

	var stu1 Student
	stu1.age = 12
	stu1.name = "张思"
	fmt.Println(stu1)
}

func stringBuild(strPifx string) func(str string) string {

	return func(str string) string {
		strPifx = strPifx + str
		return strPifx
	}
}

type Student struct {
	name string
	age  int8
}
