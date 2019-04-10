package main

import "fmt"

func main() {

	var intArr = [...]int{1, 4, 22, 11, 66}
	fmt.Println("intArr=", intArr)

	/*
		声明/定义一个切片
		1.slice切片名称
		2.intArr[1:3]表示slice引用到intArr这个数组
		3.引用intArr数组的起始下标为1，最后的下标为3（但不包括3）
	*/
	var slice = intArr[1:3]
	var slice3 = intArr[:] //表示：0:len(intArr)
	slice[0] = 99
	slice = append(slice, 10) //附加
	slice = append(slice, 20, 30, 40)

	fmt.Println("slice 元素", slice)
	fmt.Println("slice 元素个数", len(slice))
	fmt.Println("slice 容量", cap(slice))

	fmt.Println("slice3 元素", slice3)
	fmt.Println("slice3 元素个数", len(slice3))
	fmt.Println("slice3 容量", cap(slice3))

	fmt.Println("intArr=", intArr)

	slice2 := make([]float64, 5, 10)
	fmt.Println(slice2)
	slice2[0] = 1.01
	slice2 = append(slice2, 12.3) //附加
	fmt.Println(slice2)
	fmt.Printf("slice2 type=%T\n", slice2)
}
