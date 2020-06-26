package main

import "fmt"

func main2() {

	var intArray = [3]int{1, 2, 3}
	var strArray [3]string
	strArray[0] = "fuxi.hello0"
	strArray[1] = "fuxi.hello1"
	strArray[2] = "fuxi.hello2"

	fmt.Printf("intArray data type=%T\n", intArray)
	fmt.Printf("strArray value=%v\n", strArray)

	for i := 0; i < len(strArray); i++ {
		fmt.Println(strArray[i])
	}

	var ss = strArray
	ss[0] = "fuxi.hello5"

	fmt.Println("ss=", ss)
	fmt.Println("strArray=", strArray)

	slice := strArray[1:2]
	fmt.Println("slice value=", slice)
	fmt.Printf("元素个数=%d,容量=%d\n", len(slice), cap(slice))

	slice = append(slice, "slice value append1")
	slice = append(slice, "slice value append12")
	slice = append(slice, "slice value append13")
	slice = append(slice, "slice value append13")
	fmt.Println("slice value=", slice)
	fmt.Printf("元素个数=%d,容量=%d\n", len(slice), cap(slice))

	slice2 := make([]float64, 5, 8)
	slice2 = append(slice2, 12.3)
	fmt.Println("slice2 value=", slice2)
	fmt.Printf("元素个数=%d,容量=%d\n", len(slice2), cap(slice2))

	var strMap map[string]string
	strMap = make(map[string]string, 10)
	strMap["key1"] = "key1_value"

	fmt.Println("strMap=", strMap)

	strMap2 := map[string]string{"key1": "value1", "key2": "value2"}
	strMap2["key3"] = "value3"
	strMap2["key4"] = "value4"
	fmt.Println("strMap2=", strMap2)

	var funcName = func(value1 int, value2 int) int {
		return value1 + value2
	}

	fmt.Println("func return value=", funcName(12, 16))
}
