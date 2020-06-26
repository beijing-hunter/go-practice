package main

import "fmt"

func main() {

	var intArray [4][5]int

	fmt.Println("intArray length=", len(intArray))
	fmt.Println("intArray[0] length=", len(intArray[0]))

	for index := range intArray {

		for index2 := range intArray[index] {
			fmt.Printf("%v ", intArray[index][index2])
		}

		fmt.Println()
	}
}
