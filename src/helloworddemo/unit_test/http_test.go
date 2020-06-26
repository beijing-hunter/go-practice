package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHttp(t *testing.T) {
	t.Parallel() //并行运行测试用例，开头需要调用此函数
	mux := http.NewServeMux()
	mux.HandleFunc("/info", handleRequest)

	writer := httptest.NewRecorder()
	requst, _ := http.NewRequest("GET", "/info", nil)
	mux.ServeHTTP(writer, requst)

	if writer.Code != 200 {
		t.Error("请求失败")
	} else {
		var stu Student
		json.Unmarshal(writer.Body.Bytes(), &stu)
		t.Log(stu)
	}
}
