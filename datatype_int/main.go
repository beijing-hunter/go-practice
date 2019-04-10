package main

import "fmt"
import "unsafe"

func main() {

	var i int = 128
	fmt.Println("i=", i)

	var j int8 = 127
	fmt.Println("j=", j)

	fmt.Printf("j dataType=%T dataSize=%d \n", j, unsafe.Sizeof(j))
}
