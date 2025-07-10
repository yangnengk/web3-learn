package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Task func(num int)

func main() {
	towGoFunc()

	taskExecute()
}

func taskExecute() {
	/*
		设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间
	*/
	var tasks []Task

	fmt.Println(tasks)
	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		tasks = append(tasks, func(num int) {
			println("hello world ", num)
		})
	}

	for i, task := range tasks {
		go func(index int, t Task) {
			start := time.Now()
			defer wg.Done()
			num := rand.Intn(100)
			t(num)
			fmt.Println("任务task_", index, "耗时：", time.Since(start))
		}(i, task)
	}

	wg.Wait()
}

func towGoFunc() {
	/*
		编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数
	*/
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i++ {
			if i%2 == 1 {
				println("奇数：", i)
			}
		}
	}()

	go func() {
		defer wg.Done()
		for i := 2; i <= 10; i++ {
			if i%2 == 0 {
				println("偶数：", i)
			}
		}
	}()

	wg.Wait()
}
