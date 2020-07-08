package main

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"time"
)

func main()  {

	var(
		config clientv3.Config
		client *clientv3.Client
		err error
		putResponse *clientv3.PutResponse
	)

	config=clientv3.Config{
		Endpoints: []string{"127.0.0.1:2379"},
		DialTimeout: 5*time.Second,
	}

	if client,err=clientv3.New(config);err!=nil{
		fmt.Println(err)
		return
	}

	clientKv:=clientv3.NewKV(client)
	putResponse,err=clientKv.Put(context.TODO(),"cron/jobs/job1","job1")
	putResponse,err=clientKv.Put(context.TODO(),"cron/jobs/job2","job2")

	if err!=nil{
		fmt.Println(err)
		return
	}

	fmt.Println(putResponse.Header.Revision)//查看操作key的版本
}
