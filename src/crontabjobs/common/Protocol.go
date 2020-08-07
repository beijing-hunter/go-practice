package common

import (
	"encoding/json"
	"strings"
)

type Job struct {
	Name     string `json:"name"`
	Command  string `json:"command"`
	CronExpr string `json:"cronExpr"`
}

type ApiResult struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type JobEvent struct {
	EventType int //save,delete
	job       *Job
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
		job:       job,
	}

	return jobEvent
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
