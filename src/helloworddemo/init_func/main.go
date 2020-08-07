package main

import (
	"fmt"

	"init_func/utils"
)

func init() {
	fmt.Println("master.go init....")
}

func main() {

	utils.Show()
}
