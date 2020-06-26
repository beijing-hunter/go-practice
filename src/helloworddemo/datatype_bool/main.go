package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var b bool = false

	fmt.Println("b=", b)
	fmt.Printf("b dataSize=%d\n", unsafe.Sizeof(b))

}
