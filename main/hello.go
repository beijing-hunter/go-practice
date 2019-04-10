package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	fmt.Println("hello,world!")

	var strValue = "123.2353"
	var floatValue, _ = strconv.ParseFloat(strValue, 64)
	fmt.Printf("floatValue=%.2f,dataType=%T\n", floatValue, floatValue)

	var strFloat = strconv.FormatFloat(floatValue, 'f', 3, 64)
	fmt.Println(strFloat)

	var i = 12.124
	var fP *float64 = &i
	fmt.Printf("fp point address %v,fp===>%.2f\n", &fP, *fP)

	var v = int(i)
	fmt.Println(v)

	var vv = *fP + float64(v)
	fmt.Printf("%.2f\n", vv)

	fmt.Println("请输入a和b的值")
	fmt.Scanf("%d %f", &v, &vv)

	if float64(v) > vv {
		fmt.Println("v>vv")
	} else if float64(v) < vv {
		fmt.Println("v<vv")
	} else {
		fmt.Println("v==vv")
	}

	var count = 10

	for index := 0; index < count; index++ {
		fmt.Println("你好北京，我在上海。")
	}

	for {

		if count > 20 {
			fmt.Println("你好商河，我在新疆")
			break
		} else {
			fmt.Println("你好新疆，我在山西")
		}

		count++
	}

	var str = "大家快来围观，我在直播写代码"
	for index, val := range str {
		fmt.Printf("index=%v,val=%c\n", index, val)
	}

	var now = time.Now()
	fmt.Println(now.Unix())

	var values = [...]int{1, 2, 3}
	fmt.Println(values)
	fmt.Printf("values type=%T\n", values)

	var slices = make([]float64, 4)
	slices[0] = 12.23
	slices = append(slices, 14.3)
	fmt.Println(slices)
	fmt.Printf("slices type=%T\n", slices)

	for index, value := range slices {
		fmt.Println("index=", index, ",value=", value)
	}
}
