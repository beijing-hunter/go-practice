package worker

import (
	"crontabjobs/common"
	"fmt"
	"time"
)

type Scheduler struct {
	jobEventChan chan *common.JobEvent              //etcd任务事件队列
	jobPlanTable map[string]*common.JobSchedulePlan //任务调度计划表
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

//调度任务
func (scheduler *Scheduler) scheduleLoop() {
	var (
		jobEvent          *common.JobEvent
		scheduleAfterTime time.Duration
		scheduleTimer     *time.Timer
	)

	scheduleAfterTime = scheduler.TrySchedule()

	//调度的延迟定时器,NewTimer创建一个Timer，它会在最少过去时间段d后到期，向其自身的C字段发送当时的时间。
	scheduleTimer = time.NewTimer(scheduleAfterTime)

	for {
		select {
		case jobEvent = <-scheduler.jobEventChan: //监听任务变化事件
			scheduler.handlerJobEvent(jobEvent)
		case <-scheduleTimer.C: //最近的任务到期了
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
			fmt.Println("执行任务：jobName=", jobPlan.Job.Name)
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

func (scheduler *Scheduler) PushJobEvent(jobEvent *common.JobEvent) {
	scheduler.jobEventChan <- jobEvent
}

func InitScheduler() {
	G_scheduler = &Scheduler{
		jobEventChan: make(chan *common.JobEvent, 1000),
		jobPlanTable: make(map[string]*common.JobSchedulePlan),
	}

	go G_scheduler.scheduleLoop()
}
