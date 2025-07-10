package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	for i := 0; i < 10; i++ {
		//testLock()
		testAtomic()
	}
}

// 题目 ：使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值
func testAtomic() {
	var num atomic.Int64
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				num.Add(1)
			}
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("最后结果num:", num.Load())
}

// 编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
func testLock() {
	var wg sync.WaitGroup
	wg.Add(10)
	var counter = SafeCounter{Count: 0}

	for i := 0; i < 10; i++ {
		go func() {
			for n := 0; n < 1000; n++ {
				counter.Inc()
			}
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("最后结果count:", counter.getCount())
}

type SafeCounter struct {
	mu    sync.Mutex
	Count int
}

func (c *SafeCounter) Inc() {
	defer c.mu.Unlock()
	c.mu.Lock()
	c.Count++
}

func (c *SafeCounter) getCount() int {
	defer c.mu.Unlock()
	c.mu.Lock()
	return c.Count
}
