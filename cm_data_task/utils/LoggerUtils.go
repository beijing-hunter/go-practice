package utils

import (
	"io"
	"log"
	"os"
)

var (
	InfoLogger *log.Logger
)

func init() {
	infoFile, err := os.OpenFile("/Users/cmt/Desktop/work/work_code/cmtdata/cm_data_task/info.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		log.Fatalln("打开日志文件失败：", err)
	}

	InfoLogger = log.New(io.MultiWriter(os.Stderr, infoFile), "Info:", log.Ldate|log.Ltime|log.Lshortfile)
}
