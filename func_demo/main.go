package main

import "fmt"

func main() {

	var n1 = 1.1
	var n2 = 1.2
	var oper byte = '+'

	var result = operFunc(n1, n2, oper)
	fmt.Printf("result=%.2f\n", result)

	var value1, value2 = 14, 19
	value1, value2 = getSumAndSub(value1, value2) //返回多个值
	fmt.Println("value1=", value1, "value2=", value2)

	var funcName = getSumAndSub                  //函数做为变量赋值
	value1, value2 = invokeFun(10, 15, funcName) //函数做为形参传入
	fmt.Println("value1=", value1, "value2=", value2)

	sumValue := sum(10, -1, 23, 12, -9)
	fmt.Println("sumValue=", sumValue)
}

func operFunc(value1 float64, value2 float64, oper byte) (result float64) {

	switch oper {
	case '+':
		result = value1 + value2
	case '-':
		result = value1 - value2
	case '*':
		result = value1 * value2
	case '/':
		result = value1 / value2
	default:
		result = 0.0
	}

	return result
}

func getSumAndSub(value1 int, value2 int) (sum int, sub int) {
	sum = value1 + value2
	sub = value1 - value2
	return sum, sub
}

type invokeFunType func(value1 int, value2 int) (int, int) //自定义方法类型

func invokeFun(value1 int, value2 int, funcName invokeFunType) (int, int) {

	if value1 >= value2 { //一种简单的模板方法的实现
		return funcName(value1, value2)
	} else {
		var t = value1
		value1 = value2
		value2 = t
		return funcName(value1, value2)
	}
}

func sum(value int, args ...int) (result int) {

	result = value

	for index := 0; index < len(args); index++ {
		result += args[index]
	}

	return result
}
