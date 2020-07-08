package main

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"time"
)

func main() {
	var (
		config         clientv3.Config
		err            error
		client         *clientv3.Client
		kv             clientv3.KV
		getResp        *clientv3.GetResponse
		lease          clientv3.Lease
		leaseGrantResp *clientv3.LeaseGrantResponse
		putResp        *clientv3.PutResponse
	)

	config = clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	}

	if client, err = clientv3.New(config); err != nil {
		fmt.Println(err)
		return
	}

	lease = clientv3.NewLease(client)
	//申请一个10s的租约
	if leaseGrantResp, err = lease.Grant(context.TODO(), 10); err != nil {
		fmt.Println(err)
		return
	}

	//拿到租约id
	leaseId := leaseGrantResp.ID
	kv = clientv3.NewKV(client)

	if putResp, err = kv.Put(context.TODO(), "cron/lock/job1", "", clientv3.WithLease(leaseId)); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("写入成功,key版本=", putResp.Header.Revision)

	for {

		getResp, err = kv.Get(context.TODO(), "cron/lock/job1")

		if err != nil {
			fmt.Println(err)
			return
		}

		if getResp.Count == 0 {
			fmt.Println("kv过期了")
			break
		}

		fmt.Println("kv还没有过期", getResp.Kvs)
		time.Sleep(2 * time.Second)
	}
}
