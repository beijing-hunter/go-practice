package main

import (
	"fmt"
	"net/http"
)

type MyHandler struct{}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) { //实现接口方法
	fmt.Fprintf(w, "hello world!")
}

func getInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "getInfo")
}

func main() {

	myh := MyHandler{}

	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.Handle("/hello", &myh)          //处理器
	http.HandleFunc("/getinfo", getInfo) //处理器函数
	server.ListenAndServe()
}
