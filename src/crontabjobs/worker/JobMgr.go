package worker

import (
	"context"
	"crontabjobs/common"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/mvcc/mvccpb"
	"time"
)

type JobMgr struct {
	client  *clientv3.Client
	kv      clientv3.KV
	lease   clientv3.Lease
	watcher clientv3.Watcher
}

var (
	G_jobMgr *JobMgr
)

//初始化job管理
func InitJobMgr() (err error) {
	var (
		config  clientv3.Config
		client  *clientv3.Client
		kv      clientv3.KV
		lease   clientv3.Lease
		watcher clientv3.Watcher
	)

	config = clientv3.Config{
		Endpoints:   G_config.EtcdEndpoints,
		DialTimeout: time.Duration(G_config.EtcdDialTimeout) * time.Millisecond,
	}

	if client, err = clientv3.New(config); err != nil {
		return
	}

	kv = clientv3.NewKV(client)
	lease = clientv3.NewLease(client)
	watcher = clientv3.NewWatcher(client)

	G_jobMgr = &JobMgr{
		client:  client,
		kv:      kv,
		lease:   lease,
		watcher: watcher,
	}

	return G_jobMgr.WatchJobs()
}

//任务数据信息监听
func (jobMgr *JobMgr) WatchJobs() (err error) {

	var (
		getResp  *clientv3.GetResponse
		keyValue *mvccpb.KeyValue
		jobInfo  *common.Job
		jobEvent *common.JobEvent
	)

	if getResp, err = jobMgr.kv.Get(context.TODO(), common.JOB_SAVE_DIR, clientv3.WithPrefix()); err != nil {
		return err
	}

	for _, keyValue = range getResp.Kvs {

		if jobInfo, err = common.DecodJob(keyValue.Value); err != nil {
			continue
		}
		fmt.Println("jobName=", jobInfo.Name)
		jobEvent = common.BuildJobEvent(common.JOB_EVENT_SAVE, jobInfo)
		G_scheduler.PushJobEvent(jobEvent)
	}

	go func(getResp *clientv3.GetResponse) {

		var (
			watchChan         clientv3.WatchChan
			watchStartVersion int64
			jobEvent          *common.JobEvent
			jobInfo           *common.Job
			err               error
		)
		//当前etcd集群事务id,单调递增的
		watchStartVersion = getResp.Header.Revision + 1
		watchChan = jobMgr.watcher.Watch(context.TODO(), common.JOB_SAVE_DIR, clientv3.WithRev(watchStartVersion), clientv3.WithPrefix())

		//处理kv变化事件
		for watchResp := range watchChan { //watchChan没有关闭for一直执行，chan有值时执行循环，没有值时等待

			for _, event := range watchResp.Events {
				switch event.Type {
				case mvccpb.PUT:

					if jobInfo, err = common.DecodJob(event.Kv.Value); err == nil {
						jobEvent = common.BuildJobEvent(common.JOB_EVENT_SAVE, jobInfo)
						//TODO:推送job更新事件信息到任务
						G_scheduler.PushJobEvent(jobEvent)
						fmt.Println("push jobName=", jobInfo.Name)
					}

				case mvccpb.DELETE:
					jobName := common.ExtractJobName(string(event.Kv.Key))
					jobInfo = &common.Job{
						Name: jobName,
					}

					jobEvent = common.BuildJobEvent(common.JOB_EVENT_DELETE, jobInfo)
					//TODO:推送job删除事件到任务
					G_scheduler.PushJobEvent(jobEvent)
					fmt.Println("delete jobName=", jobInfo.Name)
				}
			}
		}

	}(getResp)

	return err
}

//创建任务执行lock
func (jobMgr *JobMgr) CreateJobLock(jobName string) (jobLock *JobLock) {

	jobLock = InitJobLock(jobName, jobMgr.kv, jobMgr.lease)
	return jobLock
}
