package tests

import (
	"cm_data_task/caches"
	"cm_data_task/service"
	"cm_data_task/utils"
	"fmt"
	"testing"
)

func testModuleUser(t *testing.T) {

	//defaultDataHandlerService := service.TagModuleDefaultDataHandlerService{}
	dataHandlerService := service.TagModuleDataHandlerService{}
	guessHandlerService := service.GuessLikeModuleDataHandlerService{}

	dataHandlerService.TagModuleDataCollect(true)

	guessHandlerService.Handler()
	utils.InfoLogger.Println("模块数据处理完成。。。")
}

func TestCache(t *testing.T) {

	t.Skip()
	liveCache := caches.LiveOrderfinalCache{}

	fmt.Println(liveCache.GetValue("6134201"))
}

func TestRedis(t *testing.T) {

	t.Skip() //跳过执行测试函数
	useLiveId := caches.UserUseLiveIdCache{}
	//useLiveId.AddUseLiveId(12, 13)
	fmt.Println(useLiveId.IsUseLiveId(12, 15))
}

func TestConfig(t *testing.T) {

	//t.Skip()
	iniParser := utils.IniParser{}
	iniParser.Load("../config/envs/" + utils.EnvDir + "/application.ini")
	value := iniParser.GetString("redis", "adds")
	fmt.Println(value)
}
