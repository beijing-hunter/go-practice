package main

import (
	"crontabjobs/worker"
	"flag"
	"fmt"
	"runtime"
	"time"
)

var (
	configFile string
)

func initArgs() {
	flag.StringVar(&configFile, "config", "./worker.json", "配置文件地址")
	flag.Parse()
}

func initEnv() {
	//初始化线程数量等于cpu数量
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {

	var (
		err error
	)

	//初始化命令行参数
	initArgs()

	//初始化线程
	initEnv()

	//初始化配置
	if err = worker.InitConfig(configFile); err != nil {
		goto ERR
	}

	//任务执行器
	worker.InitExecutor()

	//任务调度器
	worker.InitScheduler()

	//任务管理器
	if err = worker.InitJobMgr(); err != nil {
		goto ERR
	}

	for {
		time.Sleep(1 * time.Second)
	}

ERR:
	fmt.Println(err.Error())
}
