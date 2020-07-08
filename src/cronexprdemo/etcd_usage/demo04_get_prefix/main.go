package main

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"time"
)

func main() {
	var (
		config  clientv3.Config
		err     error
		client  *clientv3.Client
		kv      clientv3.KV
		getResp *clientv3.GetResponse
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

	//获取key的前缀等于cron/jobs/的值
	if getResp, err = kv.Get(context.TODO(), "cron/jobs/", clientv3.WithPrefix()); err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(getResp.Kvs, getResp.Count)

		for index, item := range getResp.Kvs {
			fmt.Printf("item datatype=%T,item=%v\n", item,item)
			fmt.Printf("key=%v,value=%v,index=%v\n",string(item.Key),string(item.Value),index)
		}
	}

}
