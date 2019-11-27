package main

import (
	"cm_data_task/service"
	"cm_data_task/utils"
)

func main() {

	defaultDataHandlerService := service.TagModuleDefaultDataHandlerService{}
	dataHandlerService := service.TagModuleDataHandlerService{}
	guessHandlerService := service.GuessLikeModuleDataHandlerService{}

	isCollect := dataHandlerService.IsCollectEnd()
	dataHandlerService.ClearDatas(isCollect)

	go dataHandlerService.TagModuleDataMachiningFacatory()
	dataHandlerService.TagModuleDataCollect(isCollect)

	defaultDataHandlerService.DefaultDataCollect()

	guessHandlerService.Handler()
	utils.InfoLogger.Println("模块数据处理完成。。。")
}
