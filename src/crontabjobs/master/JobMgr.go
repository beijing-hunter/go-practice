package master

import (
	"context"
	"crontabjobs/common"
	"encoding/json"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"time"
)

type JobMgr struct {
	client *clientv3.Client
	kv     clientv3.KV
	lease  clientv3.Lease
}

var (
	G_jobMgr *JobMgr
)

//初始化job管理
func InitJobMgr() (err error) {
	var (
		config clientv3.Config
		client *clientv3.Client
		kv     clientv3.KV
		lease  clientv3.Lease
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

	G_jobMgr = &JobMgr{
		client: client,
		kv:     kv,
		lease:  lease,
	}

	return
}

//保存任务信息
func (jobMgr *JobMgr) SaveJob(jobInfo common.Job) (oldJobInfo common.Job, err error) {

	var (
		jobKey   string
		jobValue []byte
		putResp  *clientv3.PutResponse
	)

	jobKey = common.JOB_SAVE_DIR + jobInfo.Name

	if jobValue, err = json.Marshal(jobInfo); err != nil {
		return oldJobInfo, err
	}

	if putResp, err = jobMgr.kv.Put(context.TODO(), jobKey, string(jobValue), clientv3.WithPrevKV()); err != nil {
		return oldJobInfo, err
	}

	if putResp.PrevKv != nil {

		if err = json.Unmarshal(putResp.PrevKv.Value, &oldJobInfo); err != nil {
			err = nil
			fmt.Println(err)
		}
	}

	return oldJobInfo, err
}

func (jobMgr *JobMgr) Delete(jobName string) (oldJobInfo common.Job, err error) {
	var (
		jobKey     string
		deleteResp *clientv3.DeleteResponse
	)

	jobKey = common.JOB_SAVE_DIR + jobName

	if deleteResp, err = jobMgr.kv.Delete(context.TODO(), jobKey, clientv3.WithPrevKV()); err != nil {
		return
	}

	if len(deleteResp.PrevKvs) != 0 {
		err = json.Unmarshal(deleteResp.PrevKvs[0].Value, &oldJobInfo)
	}

	return oldJobInfo, err
}

func (jobMgr *JobMgr) List() (datas []*common.Job, err error) {
	var (
		getResp *clientv3.GetResponse
		job     *common.Job
	)

	getResp, err = jobMgr.kv.Get(context.TODO(), common.JOB_SAVE_DIR, clientv3.WithPrefix())

	if err != nil {
		return nil, err
	}

	if len(getResp.Kvs) > 0 {

		datas = make([]*common.Job, 0)

		for _, keyItem := range getResp.Kvs {

			job = &common.Job{}

			if err = json.Unmarshal(keyItem.Value, job); err != nil {
				continue
			}

			datas = append(datas, job)
		}
	}

	return datas, err
}

//终止任务执行
func (jobMgr *JobMgr) Kill(jobName string) (err error) {
	var (
		leaseGranResp *clientv3.LeaseGrantResponse
		leaseId       clientv3.LeaseID
		killerKey     string
	)

	//获取一个1s自动过期的租约
	if leaseGranResp, err = jobMgr.lease.Grant(context.TODO(), 1); err != nil {
		return err
	}

	leaseId = leaseGranResp.ID
	killerKey = common.JOB_KILL_DIR + jobName

	if _, err = jobMgr.kv.Put(context.TODO(), killerKey, "", clientv3.WithLease(leaseId)); err != nil {
		return err
	}

	return err
}
