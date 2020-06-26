package main

import (
	"encoding/json"
	"fmt"
)

type Student struct{

	Name string `json:"name"` //小写输出
	Age int `json:"age"`
}

func jsonToStruct()  {

	jsonData:="{\"name\":\"东王太医\",\"age\":130}"
	var stu Student
	err:=json.Unmarshal([]byte(jsonData),&stu)

	if err!=nil{
		fmt.Println("反序列化失败",err)
		return
	}

	fmt.Println(stu)
}

func main()  {
	jsonToStruct()
}
