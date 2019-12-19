package main

import (
	"net/http"
	_ "net/http/pprof"
	"strings"
)

func getInfo(w http.ResponseWriter, r *http.Request) {

	var datas []string

	for i := 0; i < 10000; i++ {
		datas = append(datas, strings.ToLower("Hounter")) //演示内存消耗大的场景
	}
}

func main() {

	http.HandleFunc("/getinfo", getInfo) //处理器函数
	http.ListenAndServe("localhost:5000", http.DefaultServeMux)

	//1.GODEBUG=gctrace=1 go run main.go
	//2.go tool pprof http://localhost:5000/debug/pprof/allocs
	//3.for ((i=1;i<=1000;i++)); do curl http://localhost:5000/getinfo; done  产生gc
}
