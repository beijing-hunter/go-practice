package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main()  {

	fileInfo,error:=os.Open("F:/mnt/cmt_media/log/log.debug-2018-07-05.0.log")

	if error!=nil{
		fmt.Println("file open error,",error)
	}

	fileReader:=bufio.NewReader(fileInfo)//文件读取缓冲

	for{

		content,readError:=fileReader.ReadString('\n')//读取文件一行内容

		if readError==io.EOF{//表示，读取到文件末尾
			break
		}

		fmt.Print(content)
	}

	fileInfo.Close()
}
