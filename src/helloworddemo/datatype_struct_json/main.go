package main

import (
	"fmt"
	"encoding/json"
)

type Student struct{

	Name string `json:"name"` //小写输出
	Age int `json:"age"`
}

func main()  {
	
	stu:=Student{"李心",400}
	jsonByte,error:=json.Marshal(stu)//结构体中的字段首字母大写，才能被转换成功

	if error!=nil{
		fmt.Println("json 转换失败",error)
	}else{
		fmt.Println("jsonString=",string(jsonByte))
	}
}