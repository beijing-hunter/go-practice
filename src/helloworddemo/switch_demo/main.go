package main

import "fmt"

func main() {

	var key byte
	fmt.Println("请输入a,b,c,d,e,f,g")
	fmt.Scanf("%c", &key)

	switch key {
	case 'a':
		fmt.Println("周一")
	case 'b':
		fmt.Println("周二")
	case 'c':
		fmt.Println("周三")
	case 'd':
		fmt.Println("周四")
	case 'e':
		fmt.Println("周无")
	case 'f':
		fmt.Println("周刘")
	default:
		fmt.Println("周膜")
	}

	switch 5 {
	case 1, 3, 4, 5:
		fmt.Println("5")
	default:
		fmt.Println("hello default")
	}

}
