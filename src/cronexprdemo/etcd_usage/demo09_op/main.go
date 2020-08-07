package main

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"time"
)

func main() {

	var (
		config clientv3.Config
		err    error
		client *clientv3.Client
		kv     clientv3.KV
		op     clientv3.Op
		opResp clientv3.OpResponse
	)

	config = clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	}

	if client, err = clientv3.New(config); err != nil {
		fmt.Println(err)
		return
	}

	kv = clientv3.NewKV(client)

	//创建op
	op = clientv3.OpPut("/cron/jobs/job89", "job89")

	//执行op
	if opResp, err = kv.Do(context.TODO(), op); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("key创建版本：", opResp.Put().Header.Revision)

	op = clientv3.OpGet("/cron/jobs", clientv3.WithPrefix())

	if opResp, err = kv.Do(context.TODO(), op); err != nil {
		fmt.Println(err)
		return
	}

	for index, item := range opResp.Get().Kvs {
		fmt.Println("index=", index, "key=", string(item.Key), "value=", string(item.Value),"CreateRevision=",item.CreateRevision,"Version=",item.Version)
	}
}
