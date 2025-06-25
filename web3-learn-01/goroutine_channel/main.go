package main

import (
	"fmt"
	"time"
)

/*
goroutine 是轻量线程，创建一个 goroutine 所需的资源开销很小，所以可以创建非常多的 goroutine 来并发工作。
它们是由 Go 运行时调度的。调度过程就是 Go 运行时把 goroutine 任务分配给 CPU 执行的过程。
但是 goroutine 不是通常理解的线程，线程是操作系统调度的。
在 Go 中，想让某个任务并发或者异步执行，只需把任务封装为一个函数或闭包，交给 goroutine 执行即可。

go 中并发同样存在线程安全问题，因为 Go 也是使用共享内存让多个 goroutine 之间通信。并且大部分时候为了性能，所以 go 的大多数标准库的数据结构默认是非线程安全的
*/
func main() {
	go func() {
		fmt.Println("run goroutine in closure")
	}()

	go func(s string) {
		fmt.Println(s)
	}("goroutine: closure param")

	go say("in goroutine: world")

	say("hello")
}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}
