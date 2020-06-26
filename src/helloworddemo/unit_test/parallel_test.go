package main

import (
	"testing"
	"time"
)

func TestP_1(t *testing.T) {
	t.Parallel() //并行运行测试用例，开头需要调用此函数
	time.Sleep(1 * time.Second)
}

func TestP_2(t *testing.T) {
	t.Parallel() //并行运行测试用例，开头需要调用此函数
	time.Sleep(2 * time.Second)
}

func TestP_3(t *testing.T) {
	t.Parallel() //并行运行测试用例，开头需要调用此函数
	time.Sleep(3 * time.Second)
}
