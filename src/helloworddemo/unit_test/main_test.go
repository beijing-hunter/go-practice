package main

import "testing"

func TestSum(t *testing.T) {

	stu := Student{}
	value := stu.Sum(1, 2)

	if value > 0 {
		t.Log("输出结果正常", value)
	} else {
		t.Error("输出结果异常", value)
	}
}
