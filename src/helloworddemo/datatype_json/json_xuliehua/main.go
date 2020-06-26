package main

import (
	"encoding/json"
	"fmt"
	"github.com/json-iterator/go"
)

type Student struct{

	Name string `json:"name"` //小写输出
	Age int `json:"age"`
}

func jsonStruct()  {

	stu:=&Student{"卢布",13}

	dataByte,err:=json.Marshal(stu)

	if err!=nil{
		fmt.Println("序列化失败",err)
		return
	}

	dataStr:=string(dataByte)
	fmt.Println("json data=",dataStr)


}

func jsonMap(){

	stuMap:=make(map[string]interface{})
	stuMap["oneStu"]=&Student{"卢布",13}
	stuMap["twoStu"]=&Student{"东王太医",130}

	dataByte,err:=json.Marshal(stuMap)

	if err!=nil{
		fmt.Println("序列化失败",err)
		return
	}

	dataStr:=string(dataByte)
	fmt.Println("json data=",dataStr)
}

func jsonSlice(){

	var stuSlice []Student
	stu:=Student{"卢布",13}
	stuSlice=append(stuSlice,stu)

	stu2:=Student{"东王太医",130}
	stuSlice=append(stuSlice,stu2)

	dataByte,err:=json.Marshal(stuSlice)

	if err!=nil{
		fmt.Println("序列化失败",err)
		return
	}

	dataStr:=string(dataByte)
	fmt.Println("json data=",dataStr)
}

func main()  {
	//jsonStruct()
	//jsonMap()
	//jsonSlice()

	stuMap:=make(map[interface{}]interface{})
	stuMap["oneStu"]=&Student{"卢布",13}
	stuMap[23]=&Student{"东王太医",130}

	dataByte,err:=jsoniter.Marshal(stuMap)

	if err!=nil{
		fmt.Println("序列化失败",err)
		return
	}

	dataStr:=string(dataByte)
	fmt.Println("json data=",dataStr)
}
