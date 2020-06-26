package main

import (
	"fmt"
	"net/http"
)

func handler(respone http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(respone, "Hello World,%s!", request.URL.Path[1:])
}

func main() {

	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: nil,
	}

	http.HandleFunc("/", handler) //处理器函数
	server.ListenAndServe()
}
