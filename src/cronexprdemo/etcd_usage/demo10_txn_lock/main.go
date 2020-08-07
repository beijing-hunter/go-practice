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
		lease          clientv3.Lease
		leaseGrantResp *clientv3.LeaseGrantResponse
		txn            clientv3.Txn
		txnResp        *clientv3.TxnResponse
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
	//申请一个5s的租约
	if leaseGrantResp, err = lease.Grant(context.TODO(), 5); err != nil {
		fmt.Println(err)
		return
	}

	//拿到租约id
	leaseId := leaseGrantResp.ID
	//释放租约
	defer lease.Revoke(context.TODO(), leaseId)

	//获取锁
	kv = clientv3.NewKV(client)
	//创建事务
	txn = kv.Txn(context.TODO())

	//定义事务，if 不存在key ,then设置它，else 抢锁失败
	txn.If(clientv3.Compare(clientv3.CreateRevision("/cron/lock/job9"), "=", 0)).
		Then(clientv3.OpPut("/cron/lock/job9", "job9", clientv3.WithLease(leaseId))).
		Else(clientv3.OpGet("/cron/lock/job9"))

	if txnResp, err = txn.Commit(); err != nil {
		fmt.Println(err)
		return
	}

	//判断是否抢到了锁
	if !txnResp.Succeeded {
		fmt.Println("锁被占用了", string(txnResp.Responses[0].GetResponseRange().Kvs[0].Value))
		return
	}

	//抢到锁,处理业务逻辑
	fmt.Println("处理业务")
	time.Sleep(5 * time.Second)
}
