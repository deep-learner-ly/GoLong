package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	//var sum int
	//var wg sync.WaitGroup
	//
	//for i := 1; i <= 10000; i++ {
	//	wg.Add(1)
	//	go func(i int) {
	//		defer wg.Done()
	//		sum += i * i
	//	}(i)
	//}
	//
	//wg.Wait()
	//fmt.Println(sum)
	//
	//ch := make(chan int)
	//
	//go producer(ch)
	//go consumer(ch)

	//mq := NewMessageQueue(10)
	//
	//producer := func() {
	//	for {
	//		mq.Enqueue("msg")
	//	}
	//}
	//
	//consumer := func() {
	//	for {
	//		msg := mq.Dequeue()
	//		println(msg + "is consumed")
	//	}
	//}
	//
	//go producer()
	//go consumer()
	//
	//select {} // Block the main thread
	for i := 0; i < 1000; i++ {
		go increment()
	}

	time.Sleep(1 * time.Second)
	fmt.Println(count)
}

func producer(ch chan int) {
	for i := 1; i <= 1000; i++ {
		ch <- i
	}
	close(ch)
}

func consumer(ch chan int) {
	for i := range ch {
		fmt.Println(i)
	}
}

type MessageQueue struct {
	bufferSize int
	size       int
	queue      chan string
}

func NewMessageQueue(bufferSize int) *MessageQueue {
	return &MessageQueue{
		bufferSize: bufferSize,
		size:       0,
		queue:      make(chan string, bufferSize),
	}
}

func (mq *MessageQueue) Enqueue(message string) {
	mq.queue <- message
}

func (mq *MessageQueue) Dequeue() string {
	return <-mq.queue
}

var (
	count int
	mutex sync.Mutex
)

func increment() {
	mutex.Lock()
	count++
	mutex.Unlock()
}
