package main

import (
	"fmt"
	"os"
)

func main(){

	filePath:="F:/mnt/cmt_media/log/log.debug-2018-07-05.0.log"
	fileInfo,error:=os.Open(filePath)

	if error!=nil{
		fmt.Println("file open error ",error)
	}

	fmt.Println("file address=",fileInfo)

	error=fileInfo.Close()

	if error!=nil{
		fmt.Println("file close error",error)
	}
}
