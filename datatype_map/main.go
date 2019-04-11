package main

import "fmt"

func main() {

	//声明使用方式一
	var dataMap map[string]string
	//在使用map前，需要先用make给map分配数据空间
	dataMap = make(map[string]string, 10)
	dataMap["key1"] = "中国"

	//声明使用方式二
	dataMap2 := map[string]string{"key1": "hello", "key2": "world"}
	fmt.Println("dataMap2=", dataMap2)

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
	fmt.Println("dataMap2=", dataMap2)
}
