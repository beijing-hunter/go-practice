package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	var values [2]float64
	values[0] = 3.2

	for index, value := range values {
		fmt.Printf("index=%v,value=%v\n", index, value)
	}

	var values2 = [3]int{1, 2, 3}
	fmt.Println(values2)

	var values3 = [...]int{4, 5, 2, 1}
	fmt.Println(values3)

	var values4 = [...]int{1: 12, 0: 23, 8: 999} //指定下标赋值
	fmt.Println("values4=", values4)

	values5 := values4 //数组之间赋值，是拷贝的过程。
	values5[0] = 8
	fmt.Println("values4=", values4)
	fmt.Println("values5=", values5)

	rand.Seed(time.Now().UnixNano()) //随机数
	for i := 0; i < len(values5); i++ {
		values5[i] = rand.Intn(100)
	}

	fmt.Println("values5=", values5)
}
