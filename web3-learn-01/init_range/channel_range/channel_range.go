package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 10)

	go addDataCh(ch)

	for i := range ch { // chan 不支持第二个返回值
		fmt.Println(i)
	}
}

func addDataCh(ch chan int) {
	size := cap(ch)
	for i := 0; i < size; i++ {
		ch <- i
		time.Sleep(time.Second * 1)
	}
	close(ch) // 关闭通道,不关闭会导致死锁	fatal error: all goroutines are asleep - deadlock!
}
