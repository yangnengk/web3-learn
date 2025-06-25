package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//counter := UnsafeCounter{}	// 非安全计数器
	counter := SafeCounter{} // 安全计数器

	// 启动100个线程
	for i := 0; i < 1000; i++ {
		go func() {
			for i := 0; i < 100; i++ {
				counter.Increment()
			}
		}()
	}

	time.Sleep(time.Second * 3)

	// 最终输出计数
	fmt.Println("Final count: ", counter.count)

}

// SafeCounter 线程安全的计数器
type SafeCounter struct {
	mu    sync.Mutex
	count int
}

// Increment 增加计数
func (s *SafeCounter) Increment() {
	s.mu.Lock()
	defer s.mu.Unlock() // 方法退出的时候解锁 保证每次只有一个goroutine可以访问count
	s.count++
}

func (s *SafeCounter) GetCount() int {
	s.mu.Lock()
	defer s.mu.Unlock() // 方法退出的时候解锁
	return s.count
}

type UnsafeCounter struct {
	count int
}

// Increment 增加计数
func (c *UnsafeCounter) Increment() {
	c.count += 1
}

// GetCount 获取当前计数
func (c *UnsafeCounter) GetCount() int {
	return c.count
}
