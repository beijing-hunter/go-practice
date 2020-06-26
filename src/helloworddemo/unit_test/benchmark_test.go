package main

import (
	"fmt"
	"sync"
	"testing"
)

type Small struct {
	a int64
}

var pool = sync.Pool{
	New: func() interface{} { return new(Small) },
}

func inc(s *Small) {
	s.a++
}

// go test -v -run x -bench .
func BenchmarkWithoutPool(b *testing.B) {
	var s *Small
	for i := 0; i < b.N; i++ {
		for j := 0; j < 10000; j++ {
			s = &Small{a: 1}
			b.StopTimer()
			inc(s)
			b.StartTimer()
		}
	}
}

func BenchmarkWithPool(b *testing.B) {
	var s *Small
	for i := 0; i < b.N; i++ {
		for j := 0; j < 10000; j++ {
			s = pool.Get().(*Small)
			s.a = 1
			b.StopTimer()
			inc(s)
			b.StartTimer()
			pool.Put(s)
		}
	}
}

//性能测试
func BenchmarkCode(b *testing.B) {

	for i := 0; i < b.N; i++ {
		fmt.Println(i)
	}
}
