package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now() //获取当前时间
	fmt.Printf("now=%v\n", now)

	fmt.Printf("年=%v\n", now.Year())
	fmt.Printf("月=%v\n", int(now.Month()))
	fmt.Printf("日=%v\n", now.Day())
	fmt.Printf("时=%v\n", now.Hour())
	fmt.Printf("分=%v\n", now.Minute())
	fmt.Printf("秒=%v\n", now.Second())

	//这个字符串各个数字是固定的，必须这么写。等同于java 中的yyyy-MM-dd HH:mm:ss
	fmt.Printf("now=%v\n", now.Format("2006/01/02 15:04:05"))
	fmt.Printf("now=%v\n", now.Format("2006-01-02"))
	fmt.Printf("now=%v\n", now.Format("15:04:05"))
	fmt.Printf("now=%v\n", now.Format("01-02"))

	fmt.Println(now.Unix())     //当前时间秒时间戳
	fmt.Println(now.UnixNano()) //当前时间纳秒时间戳

	i := 0

	for {

		i++
		fmt.Println(i)
		time.Sleep(time.Second) //休眠1秒

		if i > 100 {
			break
		}
	}

}
