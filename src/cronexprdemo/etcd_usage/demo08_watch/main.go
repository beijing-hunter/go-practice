package main

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/mvcc/mvccpb"
	"time"
)

func main() {
	var (
		config            clientv3.Config
		err               error
		client            *clientv3.Client
		kv                clientv3.KV
		getResp           *clientv3.GetResponse
		watchStartVersion int64
		watcher           clientv3.Watcher
		watchChan         clientv3.WatchChan
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

	go func(kv clientv3.KV) {
		for {
			kv.Put(context.TODO(), "cron/jobs/job9", "job9")
			kv.Delete(context.TODO(), "cron/jobs/job9")
			time.Sleep(2 * time.Second)
		}
	}(kv)

	//先get到当前值，并监听后续变化
	if getResp, err = kv.Get(context.TODO(), "cron/jobs/job9"); err != nil {
		fmt.Print(err)
		return
	}

	//当前etcd集群事务id,单调递增的
	watchStartVersion = getResp.Header.Revision + 1
	//创建一个watcher
	watcher = clientv3.NewWatcher(client)
	//启动监听
	watchChan = watcher.Watch(context.TODO(), "cron/jobs/job9", clientv3.WithRev(watchStartVersion))

	//处理kv变化事件
	for watchResp := range watchChan {
		for _, event := range watchResp.Events {
			switch event.Type {
			case mvccpb.PUT:
				fmt.Println("修改为：", string(event.Kv.Value), "Revison:", event.Kv.CreateRevision, event.Kv.ModRevision)
			case mvccpb.DELETE:
				fmt.Println("删除了：", string(event.Kv.Value))
			}
		}
	}
}
