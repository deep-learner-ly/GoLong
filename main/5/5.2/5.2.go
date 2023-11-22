package main

import (
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {
	go func() {
		http.ListenAndServe("localhost:6060", nil)
	}()

	for {
		doSomeWork()
	}
}

func doSomeWork() {
	// 模拟一些工作
	time.Sleep(100 * time.Millisecond)
}
