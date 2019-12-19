package main

import (
	"fmt"
	"sort"
)

type Student struct {
	Name string
	Age  int
}

type stuSlice []Student

func (ss stuSlice) Len() int {
	return len(ss)
}

func (ss stuSlice) Less(i, j int) bool {
	return ss[i].Age > ss[j].Age
}

func (ss stuSlice) Swap(i, j int) {

	temp := ss[i]
	ss[i] = ss[j]
	ss[j] = temp
}

func main() {

	var stu stuSlice = make([]Student, 4, 5)
	stu[0] = Student{"张飞", 12}
	stu[1] = Student{"崔同山", 10}
	stu[2] = Student{"李元芳", 18}

	var stus []Student
	stus = append(stus, Student{"张飞", 12})
	stus = append(stus, Student{"阿良", 132})
	stus = append(stus, Student{"左右", 122})

	sort.SliceStable(stus, func(i int, j int) bool {
		return stus[i].Age < stus[j].Age
	})

	sort.Sort(stu)
	fmt.Println(stu)
	fmt.Println(stus)
}
