package main

import (
	"fmt"
	"net/http"
)

func getInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "getinfo")
}

func logFilter(handle http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		fmt.Println("log info", r.URL.Path)
		handle(w, r)
	}
}

func loginValidateFilter(handle http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("loginValidate 登录校验")

		if false {
			fmt.Fprintf(w, "登录失败")
		} else {
			handle(w, r)
		}
	}
}

func main() {

	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/", logFilter(loginValidateFilter(getInfo))) //串联多个处理器,可以理解为filter的一种实现方式
	server.ListenAndServe()

}
