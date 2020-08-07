package master

import (
	"crontabjobs/common"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strconv"
	"time"
)

type ApiServer struct {
	httpServer *http.Server
}

var (
	G_apiServer *ApiServer
)

//保存job信息(body json)
func handleJobSave(w http.ResponseWriter, r *http.Request) {
	var (
		bodyContent []byte
		err         error
		jobInfo     common.Job
		oldJobInfo  common.Job
		apiResult   string
	)

	//读取body体内容
	if bodyContent, err = ioutil.ReadAll(r.Body); err != nil {
		goto ERR
	}

	if err = json.Unmarshal(bodyContent, &jobInfo); err != nil {
		goto ERR
	}

	if oldJobInfo, err = G_jobMgr.SaveJob(jobInfo); err != nil {
		goto ERR
	}

	apiResult = common.BuildApiResult(0, "success", oldJobInfo)
	fmt.Fprint(w, apiResult)
	return
ERR:
	fmt.Fprint(w, common.BuildApiResult(-1, err.Error(), nil))
}

//删除任务
func handleJobDelete(w http.ResponseWriter, r *http.Request) {
	var (
		jobName    string
		err        error
		oldJobInfo common.Job
	)

	r.ParseForm()
	jobName = r.PostFormValue("jobName")
	oldJobInfo, err = G_jobMgr.Delete(jobName)

	if err != nil {
		fmt.Fprint(w, common.BuildApiResult(-1, err.Error(), nil))
		return
	}

	fmt.Fprint(w, common.BuildApiResult(0, "success", oldJobInfo))
}

func handleJobList(w http.ResponseWriter, r *http.Request) {
	var (
		jobDatas []*common.Job
		err      error
	)

	if jobDatas, err = G_jobMgr.List(); err != nil {
		fmt.Fprint(w, common.BuildApiResult(-1, err.Error(), nil))
		return
	}

	fmt.Fprint(w, common.BuildApiResult(0, "success", jobDatas))
}

func handleJobKill(w http.ResponseWriter, r *http.Request) {
	var (
		jobName string
		err     error
	)

	r.ParseForm()
	jobName = r.PostFormValue("jobName")

	if err = G_jobMgr.Kill(jobName); err != nil {
		fmt.Fprint(w, common.BuildApiResult(-1, err.Error(), nil))
		return
	}

	fmt.Fprint(w, common.BuildApiResult(0, "success", nil))
}

func InitApiServer() (err error) {
	var (
		mux        *http.ServeMux
		listener   net.Listener
		httpServer *http.Server
	)

	//配置http路由
	mux = http.NewServeMux()
	mux.HandleFunc("/job/save", handleJobSave)
	mux.HandleFunc("/job/delete", handleJobDelete)
	mux.HandleFunc("/job/list", handleJobList)
	mux.HandleFunc("/job/kill", handleJobKill)

	//web页面静态文件目录
	staticDir := http.Dir(G_config.WebRoot)
	staticHandler := http.FileServer(staticDir)
	mux.Handle("/", http.StripPrefix("/", staticHandler))

	//启动tcp监听
	if listener, err = net.Listen("tcp", "127.0.0.1:"+strconv.Itoa(G_config.ApiPort)); err != nil {
		return
	}

	//配置http服务
	httpServer = &http.Server{
		ReadTimeout:  time.Duration(G_config.ApiReadTimeout) * time.Millisecond,
		WriteTimeout: time.Duration(G_config.ApiWriteTimeout) * time.Millisecond,
		Handler:      mux,
	}

	G_apiServer = &ApiServer{
		httpServer: httpServer,
	}

	go httpServer.Serve(listener)
	return
}
