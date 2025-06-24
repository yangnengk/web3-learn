package main

import (
	"fmt"
	"sync"
	"time"
)

// map作为参数,
// map 集合也是引用类型，和切片一样，将 map 集合作为参数传给函数或者赋值给另一个变量，它们都指向同一个底层数据结构，对 map 集合的修改，都会影响到原始实参。
// 扩容也会影响到原始实参
/*
运行 main()函数，会直接报错提示：
fatal error: concurrent map writes
如果读写操作同时存在也会报错提示：
fatal error: concurrent map read and map write
当 map 集合会被并发访问时，需要在使用 map 集合时，添加互斥锁：
	也可以使用 Go 标准库中的实现 sync.Map，但是 sync.Map 适用于读多写少的场景，并且内存开销会比普通的 map 集合更大。
	所以碰到这种情况，更推荐使用普通的互斥锁来保证 map 集合的并发读写的线程安全性。
*/
func main() {
	// m := make(map[string]int, 2)
	// m["hello"] = 100

	// receiveMap(m)
	// fmt.Println("in main, after modify param:", m)

	// 并发时使用map
	var wg sync.WaitGroup
	var lock sync.Mutex
	wg.Add(2)

	m := make(map[string]int)

	go func() {
		for {
			lock.Lock()
			m["a"]++
			lock.Unlock()
		}
	}()

	go func() {
		for {
			lock.Lock()
			m["a"]++
			fmt.Println(m["a"])
			lock.Unlock()
		}
	}()

	select {
	case <-time.After(time.Second * 5):
		fmt.Println("time out, stop")
	}
}

func receiveMap(param map[string]int) {
	fmt.Println("in func, before modify param:", param)
	param["hello"] = 200
	param["world"] = 300
	param["test"] = 400
	param["test1"] = 500
}
