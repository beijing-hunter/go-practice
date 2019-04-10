package main

import "fmt"

func main() {

	var s string = "hello world "
	fmt.Println(s)

	ss := s + " \nbei jing"
	fmt.Println(ss) //转义字符识别，输出

	var ssss = `hello \n bei jing`
	fmt.Println(ssss) //不识别转义字符，原生输出

	//当一个拼接的操作很长时，可以分行写,但是+号需要留在上一行
	var s4 string = "hello wold" + "hello wold" + "hello wold" +
		"hello wold" + "hello wold" + "hello wold" +
		"hello wold" + "hello wold" + "hello wold"
	fmt.Println(s4)

	fmt.Println("s4 length ", len(s4))
}
