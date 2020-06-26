package main

import "fmt"

/**
切片与map的组合使用
**/
func main() {

	sliceMap := make([]map[string]string, 3)
	sliceMap[0] = make(map[string]string)
	sliceMap[0]["name"] = "zhangsan"

	sliceMap[1] = map[string]string{"name": "wangsan"}

	map1 := map[string]string{"name": "lisan"}
	sliceMap = append(sliceMap, map1)
	map1["sex"] = "100"

	fmt.Println("sliceMap=", sliceMap)

}
