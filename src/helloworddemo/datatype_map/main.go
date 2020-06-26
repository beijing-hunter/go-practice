package main

import (
	"fmt"
	"sort"
)

func main() {

	//声明使用方式一
	var dataMap map[string]string
	//在使用map前，需要先用make给map分配数据空间
	dataMap = make(map[string]string, 10)
	dataMap["key1"] = "中国"

	//声明使用方式二
	dataMap2 := map[string]string{"key1": "hello", "bkey2": "world", "akey2": "world"}
	dataMap2["ekey"] = "sse"
	fmt.Println("排序之前dataMap2=", dataMap2)

	stuMap := make(map[string]map[string]string)
	stuMap["01"] = make(map[string]string, 2)
	stuMap["01"]["realName"] = "张思"
	stuMap["01"]["address"] = "美国"

	stuMap["02"] = make(map[string]string, 2)
	stuMap["02"]["realName"] = "张思女"
	stuMap["02"]["address"] = "南非"

	keyName := "01"
	fmt.Println(stuMap[keyName])

	delete(dataMap2, "key1") //删除操作
	fmt.Println("删除之后dataMap2=", dataMap2)

	val, isReadSuccess := stuMap["01"] //查找map
	if isReadSuccess {
		fmt.Println(val["realName"])
	} else {
		fmt.Println("key不存在")
	}

	for key, value := range stuMap {

		fmt.Printf("%v=%v\n", key, value)

		for key2, value2 := range value {
			fmt.Printf("%v=%v ", key2, value2)
		}

		fmt.Println()
	}

	dataIntMap := map[int]int{4: 12, 2: 90, 9: 12, 5: 19} //map默认是进行了递增排序。
	fmt.Println("排序之前dataIntMap=", dataIntMap)

	dataMap2 = mapOrderBy(dataMap2)
	fmt.Println("排序之后dataMap2=", dataMap2)
}

func mapOrderBy(dataMap map[string]string) (resultMap map[string]string) {

	keys := make([]string, len(dataMap))

	i := 0
	for key := range dataMap {
		keys[i] = key
		i++
	}

	sort.Strings(keys)
	resultMap = make(map[string]string, len(dataMap))

	for _, value := range keys {
		resultMap[value] = dataMap[value]
	}

	return resultMap
}
