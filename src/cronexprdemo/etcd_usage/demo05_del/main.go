package main

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
)

func main() {
	var (
		err     error
		client  *clientv3.Client
		delResp *clientv3.DeleteResponse
	)

	config := clientv3.Config{
		Endpoints: []string{"127.0.0.1:2379"},
	}

	client, err = clientv3.New(config)

	if err != nil {
		fmt.Println(err)
		return
	}

	kv := clientv3.NewKV(client)
	//获取key操作之前的vale,clientv3.WithPrevKV()
	delResp, err = kv.Delete(context.TODO(), "cron/jobs/job1", clientv3.WithPrevKV())

	if err != nil {
		fmt.Println(err)
		return
	}

	if len(delResp.PrevKvs) > 0 { //只有加上clientv3.WithPrevKV()选项才会有值

		for _, kvItem := range delResp.PrevKvs {
			fmt.Println("删除了", string(kvItem.Key), string(kvItem.Value))
		}
	}
}
