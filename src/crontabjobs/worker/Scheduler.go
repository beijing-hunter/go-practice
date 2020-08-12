package worker

import (
	"crontabjobs/common"
	"fmt"
	"time"
)

type Scheduler struct {
	jobEventChan         chan *common.JobEvent              //etcd任务事件队列
	jobPlanTable         map[string]*common.JobSchedulePlan //任务调度计划表
	jobExecutingTable    map[string]*common.JobExecuteInfo  //任务执行表
	jobExecuteResultChan chan *common.JobExecuteResult
}

var (
	G_scheduler *Scheduler
)

//处理任务事件
func (scheduler *Scheduler) handlerJobEvent(jobEvent *common.JobEvent) {
	var (
		jobSchedulerPlan *common.JobSchedulePlan
		err              error
		jobIsExisted     bool
	)
	switch jobEvent.EventType {
	case common.JOB_EVENT_SAVE:

		if jobSchedulerPlan, err = common.BuildSchedulerPlan(jobEvent.Job); err != nil {
			return
		}

		scheduler.jobPlanTable[jobEvent.Job.Name] = jobSchedulerPlan
	case common.JOB_EVENT_DELETE:

		if jobSchedulerPlan, jobIsExisted = scheduler.jobPlanTable[jobEvent.Job.Name]; jobIsExisted {
			delete(scheduler.jobPlanTable, jobEvent.Job.Name)
		}
	}
}

//处理任务执行结果
func (scheduler *Scheduler) handlerJobResult(result *common.JobExecuteResult) {

	delete(scheduler.jobExecutingTable, result.JobExecuteInfo.Job.Name)
	fmt.Println("任务执行完成", result.JobExecuteInfo.Job.Name,string(result.OutPut),result.Err)
}

//调度协程
func (scheduler *Scheduler) scheduleLoop() {
	var (
		jobEvent          *common.JobEvent
		scheduleAfterTime time.Duration
		scheduleTimer     *time.Timer
		jobExecuteResult  *common.JobExecuteResult
	)

	scheduleAfterTime = scheduler.TrySchedule()

	//调度的延迟定时器,NewTimer创建一个Timer，它会在最少过去时间段d后到期，向其自身的C字段发送当时的时间。
	scheduleTimer = time.NewTimer(scheduleAfterTime)

	for {
		select {
		case jobEvent = <-scheduler.jobEventChan: //监听任务变化事件
			scheduler.handlerJobEvent(jobEvent)
		case <-scheduleTimer.C: //最近的任务到期了
		case jobExecuteResult = <-scheduler.jobExecuteResultChan: //任务执行结果
			scheduler.handlerJobResult(jobExecuteResult)
		}

		//调度一次任务
		scheduleAfterTime = scheduler.TrySchedule()
		//重置调度时间
		scheduleTimer.Reset(scheduleAfterTime)
	}
}

//执行任务
func (scheduler *Scheduler) TrySchedule() (scheduleAfterTime time.Duration) {

	var (
		now      time.Time
		jobPlan  *common.JobSchedulePlan
		nearTime *time.Time
	)

	if len(scheduler.jobPlanTable) == 0 {
		scheduleAfterTime = 1 * time.Second
		return
	}

	now = time.Now()

	for _, jobPlan = range scheduler.jobPlanTable {

		if jobPlan.NextTime.Before(now) || jobPlan.NextTime.Equal(now) {
			//TODO:执行任务
			scheduler.tryStartJob(jobPlan)
			jobPlan.NextTime = jobPlan.Expr.Next(now) //设置下次执行时间
		}

		//统计最近一个要过期的任务时间（用于计算调度休眠时间）
		if nearTime == nil || jobPlan.NextTime.Before(*nearTime) {
			nearTime = &jobPlan.NextTime
		}
	}

	//下次调度间隔（最近要执行的任务调度-当前时间）
	scheduleAfterTime = (*nearTime).Sub(now)
	return scheduleAfterTime
}

//尝试执行任务
func (scheduler *Scheduler) tryStartJob(plan *common.JobSchedulePlan) {
	var (
		jobExecuteInfo *common.JobExecuteInfo
		jobExecuting   bool
	)

	if jobExecuteInfo, jobExecuting = scheduler.jobExecutingTable[plan.Job.Name]; jobExecuting {
		fmt.Println("任务正在执行，跳过执行jobName=", plan.Job.Name)
		return
	}

	jobExecuteInfo = common.BuildJobExecuteInfo(plan)

	//保存任务执行信息
	scheduler.jobExecutingTable[plan.Job.Name] = jobExecuteInfo

	//执行任务
	fmt.Println("执行任务jobName=", jobExecuteInfo.Job.Name, jobExecuteInfo.PlanTime, jobExecuteInfo.RealTime)
	G_executor.ExecuteJob(jobExecuteInfo)
}

func (scheduler *Scheduler) PushJobEvent(jobEvent *common.JobEvent) {
	scheduler.jobEventChan <- jobEvent
}

//写入任务执行结果
func (scheduler *Scheduler) PushJobExecuteResult(result *common.JobExecuteResult) {
	scheduler.jobExecuteResultChan <- result
}

func InitScheduler() {
	G_scheduler = &Scheduler{
		jobEventChan:         make(chan *common.JobEvent, 1000),
		jobPlanTable:         make(map[string]*common.JobSchedulePlan),
		jobExecutingTable:    make(map[string]*common.JobExecuteInfo),
		jobExecuteResultChan: make(chan *common.JobExecuteResult, 1000),
	}

	go G_scheduler.scheduleLoop()
}
