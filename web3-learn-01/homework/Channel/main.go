package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//sendAndReceive()

	productAndConsumer()
}

// 生产者消费者模型,实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印
func productAndConsumer() {
	ch := make(chan int, 20)

	var wg sync.WaitGroup
	wg.Add(2)

	go productor(ch, &wg)
	go consumer(ch, &wg)

	wg.Wait()

}

func productor(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(ch)
	for i := 0; i < 100; i++ {
		ch <- i
		fmt.Println("发送数据：", i)
	}
}

func consumer(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	timeout := time.After(10 * time.Second)
	for {
		select {
		case v, ok := <-ch:
			if !ok {
				fmt.Println("通道已关闭")
				return
			}
			fmt.Println("接收数据：", v)
		case <-timeout:
			fmt.Println("超时")
		default:
			fmt.Println("没有接收数据，休息一会")
		}
	}
}

func sendAndReceive() {
	// 题目 ：编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来
	var wg sync.WaitGroup
	wg.Add(2)

	ch := make(chan int)
	go send(ch, &wg)
	go receive(ch, &wg)

	wg.Wait()
}

// 发送数据
func send(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(ch) // 关闭通道
	for i := 0; i < 10; i++ {
		ch <- i
		//fmt.Println("发送数据：", i)
	}
}

// 接收数据
func receive(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	//for c := range ch {
	//	fmt.Println("接收数据：", c)
	//}
	timeout := time.After(10 * time.Second)
	for {
		select {
		case c, ok := <-ch:
			if !ok {
				fmt.Println("通道已关闭")
				return
			}
			fmt.Println("接收数据：", c)
		case <-timeout:
			fmt.Println("超时")
			return
		default:
			fmt.Println("没有接收数据，休息一会")
		}
	}
}
