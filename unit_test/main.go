package main

import (
	"encoding/json"
	"net/http"
)

type Student struct {
	Name string
	Age  int8
}

func (stu Student) Sum(n int64, m int64) int64 {

	return n + m
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	stu := Student{
		Name: "年底了",
		Age:  12,
	}

	json, _ := json.Marshal(&stu)
	w.Write(json)
}

func main() {

}
