package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	pool = sync.Pool{
		New: func() interface{} { return new(Student) },
	}
)

type Student struct {
	Name string
}

func setName() {
	s := pool.Get().(*Student)
	s.Name = "shubaofei"
}

func main() {

	stu := pool.Get().(*Student) //执行步骤：1.从池中查询是否有值，有值则pop出值，无值则调用New函数创建一个student。
	stu.Name = "hello world"
	fmt.Printf("stu address=%p,value=%v\n", stu, stu)
	pool.Put(stu)
	//1.将值存入池中，但是值不是一直存放在池中的。
	//go 1.12：gc会直接回收调池所占用的内存；
	//go 1.13:gc回收时，新策略非常简单。现在有两组池：活动池和存档池（译者注：allPools 和 oldPools）。
	//当 GC 运行时，它会将每个池的引用保存到池中的新属性（victim），然后在清理当前池之前将该组池变成存档池：

	go setName()
	time.Sleep(time.Second * 10) //休眠10秒

	stu2 := pool.Get().(*Student)
	fmt.Printf("stu2 address=%p,value=%v\n", stu2, stu2)

}
