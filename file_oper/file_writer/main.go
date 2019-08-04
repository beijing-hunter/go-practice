package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){

	/**
	  O_RDONLY int = syscall.O_RDONLY // 只读模式打开文件
	    O_WRONLY int = syscall.O_WRONLY // 只写模式打开文件
	    O_RDWR   int = syscall.O_RDWR   // 读写模式打开文件
	    O_APPEND int = syscall.O_APPEND // 写操作时将数据附加到文件尾部
	    O_CREATE int = syscall.O_CREAT  // 如果不存在将创建一个新文件
	    O_EXCL   int = syscall.O_EXCL   // 和O_CREATE配合使用，文件必须不存在
	    O_SYNC   int = syscall.O_SYNC   // 打开文件用于同步I/O
	    O_TRUNC  int = syscall.O_TRUNC  // 如果可能，打开时清空文件
	 */
	filePath:="f:/go_hello.txt"
	fileInfo,err:=os.OpenFile(filePath,os.O_APPEND|os.O_CREATE,0666)

	if err !=nil{
		fmt.Println("file open error,",err)
	}

	defer fileInfo.Close()

	conent:="hello beijing\n"
	fileWriter:=bufio.NewWriter(fileInfo)

	for i:=0;i<5;i++{
		fileWriter.WriteString(conent)//写入内容到缓冲区
	}

	fileWriter.Flush()//刷新缓冲区
}
