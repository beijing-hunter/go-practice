package worker

import (
	"context"
	"crontabjobs/common"
	"os/exec"
	"time"
)

type Executor struct {
}

var (
	G_executor *Executor
)

//执行一个任务
func (executor *Executor) ExecuteJob(info *common.JobExecuteInfo) {
	go func() {
		var (
			cmd     *exec.Cmd
			output  []byte
			err     error
			result  *common.JobExecuteResult
			jobLock *JobLock
		)

		result = &common.JobExecuteResult{
			JobExecuteInfo: info,
			OutPut:         make([]byte, 0),
		}

		jobLock = G_jobMgr.CreateJobLock(info.Job.Name)
		err = jobLock.TryLock()
		defer jobLock.UnLock()

		if jobLock.IsLockSuccess { //获取锁成功

			result.StartTime = time.Now()
			//执行shell脚本
			cmd = exec.CommandContext(context.TODO(), "/bin/bash", "-c", info.Job.Command)
			output, err = cmd.CombinedOutput()

			result.EndTime = time.Now()
			result.OutPut = output
			result.Err = err
		} else {
			result.Err = err
		}

		G_scheduler.PushJobExecuteResult(result)
	}()
}

func InitExecutor() {
	G_executor = &Executor{}
}
