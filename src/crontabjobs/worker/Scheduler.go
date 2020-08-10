package worker

import "crontabjobs/common"

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
	var jobEvent *common.JobEvent

	for {
		select {
		case jobEvent = <-scheduler.jobEventChan: //监听任务变化事件
			scheduler.handlerJobEvent(jobEvent)
		}
	}
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
