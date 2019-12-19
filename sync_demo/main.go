package main

import (
	"fmt"
	"sync"
)

func main() {
	var smap sync.Map //协程安全的map
	smap.Store("name", "ceshi")
	smap.Store("name1", "ceshi1")
	smap.Store("name2", "ceshi2")
	smap.Store("name3", "ceshi3")

	value, ok := smap.Load("name4")
	fmt.Println(value, ok)

	smap.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		return true
	})
}
