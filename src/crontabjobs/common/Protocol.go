package common

import (
	"encoding/json"
	"github.com/gorhill/cronexpr"
	"strings"
	"time"
)

type Job struct {
	Name     string `json:"name"`
	Command  string `json:"command"`
	CronExpr string `json:"cronExpr"`
}

//任务调度计划
type JobSchedulePlan struct {
	Job      *Job
	Expr     *cronexpr.Expression //解析好的cronexpr表达式
	NextTime time.Time
}

type ApiResult struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type JobEvent struct {
	EventType int //save,delete
	Job       *Job
}

//构建Api响应结果
func BuildApiResult(code int, msg string, data interface{}) (jsonValue string) {

	apiResult := ApiResult{
		Code: code,
		Msg:  msg,
		Data: data,
	}

	jsonValueByte, _ := json.Marshal(apiResult)
	jsonValue = string(jsonValueByte)
	return jsonValue
}

func BuildJobEvent(eventType int, job *Job) (jobEvent *JobEvent) {
	jobEvent = &JobEvent{
		EventType: eventType,
		Job:       job,
	}

	return jobEvent
}

//构建调度计划
func BuildSchedulerPlan(job *Job) (plan *JobSchedulePlan, err error) {
	var (
		expr *cronexpr.Expression
	)

	if expr, err = cronexpr.Parse(job.CronExpr); err != nil {
		return
	}

	plan = &JobSchedulePlan{
		Job:      job,
		Expr:     expr,
		NextTime: expr.Next(time.Now()),
	}

	return plan, err
}

func DecodJob(jobValue []byte) (job *Job, err error) {
	job = &Job{}
	err = json.Unmarshal(jobValue, job)
	return job, err
}

func ExtractJobName(jobKey string) (jobName string) {
	jobName = strings.TrimPrefix(jobKey, JOB_SAVE_DIR)
	return jobName
}
