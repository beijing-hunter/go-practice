package main

import (
	"fmt"
	"github.com/gorhill/cronexpr"
	"time"
)

type CronJob struct {
	expr *cronexpr.Expression
	nextTime time.Time
	exec execFuc
}

type execFuc func()

func main()  {

	var(
		now time.Time
		expr *cronexpr.Expression
		scheduleTable map[string]*CronJob
		cronJob *CronJob
	)

	now=time.Now()
	scheduleTable=make(map[string]*CronJob)

	//5毫秒
	expr=cronexpr.MustParse("*/5 * * * * * *")
	cronJob=&CronJob{
		expr:expr,
		nextTime:expr.Next(now),
		exec: func() {
			fmt.Print("执行job1")
		},
	}

	//注册任务
	scheduleTable["job1"]=cronJob

	expr=cronexpr.MustParse("*/6 * * * * * *")
	cronJob=&CronJob{
		expr:expr,
		nextTime:expr.Next(now),
		exec: func() {
			fmt.Println("执行job2")
		},
	}

	scheduleTable["job2"]=cronJob

	go func() {
		for {

			var (
				cronJob *CronJob
				now time.Time
			)
			now=time.Now()
			for _,cronJob=range scheduleTable{

				if cronJob.nextTime.Before(now)||cronJob.nextTime.Equal(now){
					go cronJob.exec()
					cronJob.nextTime=cronJob.expr.Next(now)
				}
			}

			select {
			case <-time.NewTimer(100*time.Millisecond).C://将在100毫秒可读，返回
				
			}
		}
	}()

	time.Sleep(60*time.Second)
}
