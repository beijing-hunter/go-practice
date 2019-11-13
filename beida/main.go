package main

import (
	"beida/service"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

var (
	userService   = service.UserService{}
	schoolService = service.SchoolService{}
	loginMap      map[string]service.User
)

func init() {
	loginMap = make(map[string]service.User)
}

type ApiResult struct {
	StatusCode int64
	Result     interface{}
}

func login(w http.ResponseWriter, r *http.Request) {

	name := r.FormValue("name")
	password := r.FormValue("password")
	datas := userService.Login(name, password)

	if len(datas) == 0 {
		apiresult := ApiResult{StatusCode: 1, Result: "用户名或密码不正确"}
		jsonByte, _ := json.Marshal(apiresult)
		fmt.Fprintf(w, string(jsonByte))
		return
	}

	loginMap[datas[0].Uuid] = datas[0]
	apiresult := ApiResult{StatusCode: 0, Result: datas[0].Uuid}
	jsonByte, _ := json.Marshal(apiresult)

	setJSONRespone(jsonByte, &w)
}

func addUser(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	password := r.FormValue("password")
	info := service.User{Name: name, Password: password}

	result := userService.AddInfo(info)
	apiresult := ApiResult{}

	if result {
		apiresult = ApiResult{StatusCode: 0, Result: "添加成功"}
	} else {
		apiresult = ApiResult{StatusCode: 1, Result: "用户名称重复"}
	}

	jsonByte, _ := json.Marshal(apiresult)
	setJSONRespone(jsonByte, &w)
}

func updatePassword(w http.ResponseWriter, r *http.Request) {

	uuid := r.Header.Get("token")
	password := r.FormValue("password")
	userService.UpdatePassword(uuid, password)

	apiresult := ApiResult{StatusCode: 0, Result: "修改成功"}
	jsonByte, _ := json.Marshal(apiresult)
	setJSONRespone(jsonByte, &w)
}

func findSchoolInfos(w http.ResponseWriter, r *http.Request) {

	searchKey := r.FormValue("searchKey")
	area := r.FormValue("area")
	pageIndex := r.FormValue("pageIndex")
	pageSize := r.FormValue("pageSize")
	pageN, _ := strconv.ParseInt(pageIndex, 10, 64)
	pageS, _ := strconv.ParseInt(pageSize, 10, 64)
	datas := schoolService.FindInfos(searchKey, area, pageN, pageS)

	apiresult := ApiResult{StatusCode: 0, Result: datas}
	jsonByte, _ := json.Marshal(apiresult)
	setJSONRespone(jsonByte, &w)
}

func addSchool(w http.ResponseWriter, r *http.Request) {

	area := r.FormValue("area")
	name := r.FormValue("name")

	info := service.School{Name: name, Area: area}
	result := schoolService.AddInfo(info)

	apiresult := ApiResult{}

	if result {
		apiresult = ApiResult{StatusCode: 0, Result: "添加成功"}
	} else {
		apiresult = ApiResult{StatusCode: 1, Result: "学校名称重复"}
	}

	jsonByte, _ := json.Marshal(apiresult)
	setJSONRespone(jsonByte, &w)
}

func delSchool(w http.ResponseWriter, r *http.Request) {
	idStr := r.FormValue("id")
	id, _ := strconv.ParseInt(idStr, 10, 64)
	schoolService.Del(id)

	apiresult := ApiResult{StatusCode: 0, Result: "删除成功"}
	jsonByte, _ := json.Marshal(apiresult)
	setJSONRespone(jsonByte, &w)
}

func setJSONRespone(dataByte []byte, w *http.ResponseWriter) {

	(*w).Header().Set("Content-Type", "application/json")
	(*w).Write(dataByte)
}

func loginValidateFilter(handle http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		uuid := r.Header.Get("token")
		apiresult := ApiResult{StatusCode: 1, Result: "未登录"}
		jsonByte, _ := json.Marshal(apiresult)

		if len(uuid) == 0 {
			fmt.Fprintf(w, string(jsonByte))
			return
		}

		_, isSuccess := loginMap[uuid]

		if isSuccess { //已登录
			handle(w, r)
		} else {
			fmt.Fprintf(w, string(jsonByte))
		}
	}
}

func main() {

	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/user/login", login)
	http.HandleFunc("/user/add", addUser)
	http.HandleFunc("/user/up", updatePassword)

	http.HandleFunc("/school/finds", findSchoolInfos)
	http.HandleFunc("/school/add", loginValidateFilter(addSchool))
	http.HandleFunc("/school/del", loginValidateFilter(delSchool))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello world!")
	})

	server.ListenAndServe()

}
