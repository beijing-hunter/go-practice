package worker

import (
	"context"
	"crontabjobs/common"
	"fmt"
	"github.com/coreos/etcd/clientv3"
)

type JobLock struct {
	kv            clientv3.KV
	lease         clientv3.Lease
	jobName       string
	IsLockSuccess bool
	leaseId       clientv3.LeaseID
}

func InitJobLock(jobName string, kv clientv3.KV, lease clientv3.Lease) (jobLock *JobLock) {
	jobLock = &JobLock{
		kv:            kv,
		lease:         lease,
		jobName:       jobName,
		IsLockSuccess: false,
	}

	return jobLock
}

//获取锁
func (jobLock *JobLock) TryLock() (err error) {
	var (
		leaseGrantResp *clientv3.LeaseGrantResponse
		txn            clientv3.Txn
		txnResp        *clientv3.TxnResponse
		lockKey        string
	)

	//申请一个5s的租约
	if leaseGrantResp, err = jobLock.lease.Grant(context.TODO(), 5); err != nil {
		fmt.Println(err)
		return
	}

	//拿到租约id
	leaseId := leaseGrantResp.ID

	//创建事务
	txn = jobLock.kv.Txn(context.TODO())

	lockKey = common.JOB_LOCK_DIR + jobLock.jobName
	//定义事务，if 不存在key ,then设置它，else 抢锁失败
	txn.If(clientv3.Compare(clientv3.CreateRevision(lockKey), "=", 0)).
		Then(clientv3.OpPut(lockKey, "", clientv3.WithLease(leaseId))).
		Else(clientv3.OpGet(lockKey))

	if txnResp, err = txn.Commit(); err != nil {
		goto FAIL
	}

	//判断是否抢到了锁
	if !txnResp.Succeeded {
		err = common.JOB_LOCK_ERROR
		goto FAIL
	}

	jobLock.leaseId = leaseId
	jobLock.IsLockSuccess = true
	return err

FAIL:
	//释放租约
	jobLock.lease.Revoke(context.TODO(), leaseId)
	return err
}

//释放锁
func (jobLock *JobLock) UnLock() {

	if jobLock.IsLockSuccess {
		jobLock.lease.Revoke(context.TODO(), jobLock.leaseId)
	}
}
