package main

import (
	"fmt"
	"github.com/gorhill/cronexpr"
	"time"
)

func main()  {

	var (
		expr *cronexpr.Expression
		err error
		now time.Time
	)

	//每1分钟执行
	if expr,err=cronexpr.Parse("* * * * *");err!=nil{
		fmt.Println(err)
		return
	}

	//每5分钟执行1次
	if expr,err=cronexpr.Parse("*/5 * * * *");err!=nil{
		fmt.Println(err)
		return
	}

	//每1毫秒执行1次
	if expr,err=cronexpr.Parse("* * * * * * *");err!=nil{
		fmt.Println(err)
		return
	}

	now=time.Now()
	//下次调度时间
	nextTime:=expr.Next(now)
	fmt.Println(now,nextTime)

	time.AfterFunc(nextTime.Sub(now), func() {
		fmt.Println("被调度了：",nextTime)
	})

	time.Sleep(5*time.Second)

}
