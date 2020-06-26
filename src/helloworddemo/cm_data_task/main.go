package main

import (
	"cm_data_task/service"
	"cm_data_task/utils"
	"flag"
)

var (
	handlerFuncName = ""
)

func init() {

	flag.StringVar(&handlerFuncName, "handlerFuncName", "modelDataHandler", "执行那个数据处理函数")
}

func main() { //main函数启动命令如：go run main.go -handlerFuncName moduleDataHandler

	flag.Parse() //暂停获取参数

	switch handlerFuncName {
	case "moduleDataHandler":
		moduleDataHandler()
	default:
		utils.InfoLogger.Println("没有匹配到要执行的数据处理函数")
	}

}

//模块数据处理函数
func moduleDataHandler() {

	defaultDataHandlerService := service.TagModuleDefaultDataHandlerService{}
	dataHandlerService := service.TagModuleDataHandlerService{}
	guessHandlerService := service.GuessLikeModuleDataHandlerService{}

	dataHandlerService.ClearDatas(true)

	go dataHandlerService.TagModuleDataMachiningFacatory()
	dataHandlerService.TagModuleDataCollect(true)

	defaultDataHandlerService.DefaultDataCollect()

	guessHandlerService.Handler()
	utils.InfoLogger.Println("模块数据处理完成。。。")
}
