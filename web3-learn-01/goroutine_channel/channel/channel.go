package main

import (
	"fmt"
	"time"
)

/*
channel 是 Go 中定义的一种类型，专门用来在多个 goroutine 之间通信的线程安全的数据结构。
可以在一个 goroutine 中向一个 channel 中发送数据，从另外一个 goroutine 中接收数据。
channel 类似队列，满足先进先出原则。
channel 关闭后，就不能再发送数据了，但是可以继续接收数据，接收完数据后，ok_flag 为 false，并且接受的数据为数据结构的默认值。
<------------------------------------------------------------------------------------------------------------------->
在 Go 中，当需要 goroutine 之间协作地方，更常见的方式是使用 channel，而不是 sync 包中的 Mutex 或 RWMutex 的互斥锁。但其实它们各有侧重。
大部分时候，流程是根据数据驱动的，channel 会被使用得更频繁。
channel 擅长的是数据流动的场景：
1.传递数据的所有权，即把某个数据发送给其他协程。
2.分发任务，每个任务都是一个数据。
3.交流异步结果，结果是一个数据。
而锁使用的场景更偏向同一时间只给一个协程访问数据的权限：
1.访问缓存
2.管理状态
*/
func main() {
	/*
		仅声明
		var <channel_name> chan <type_name>
		初始化
		<channel_name> := make(chan <type_name>)
		初始化有缓冲的channel
		<channel_name> := make(chan <type_name>, 3)
	*/

	/*
		发送数据
		channel_name <- variable_name_or_value
		接收数据
		value_name, ok_flag := <- channel_name
		value_name := <- channel_name
		关闭channel
		close(channel_name)
	*/

	// channel 还有两个变种，可以把 channel 作为参数传递时，限制 channel 在函数或方法中能够执行的操作
	/*
		仅发送数据
		func <method_name>(<channel_name> chan <- <type>)
		仅接收数据
		func <method_name>(<channel_name> <-chan <type>)
	*/

	// 创建一个带缓冲的channel（缓存为3），超过3个数据会阻塞, 没有缓存默认缓存是1
	//ch := make(chan int, 3)
	ch := make(chan int)

	// 启动 发送goroutine
	go sendOnly(ch)

	// 启动 接收goroutine
	//go receiveOnly(ch)

	// 使用select进行多路复用
	timeout := time.After(10 * time.Second) // timeout实际也是一个channel
	for {
		select {
		case v, ok := <-ch:
			if !ok {
				fmt.Println("channel closed")
				return
			}
			time.Sleep(500 * time.Millisecond)
			fmt.Println("主goroutine 接收到数据：", v)
		case <-timeout:
			fmt.Println("操作超时")
			return
		default:
			fmt.Println("没有接收到数据，等待中...")
			time.Sleep(500 * time.Millisecond)
		}
	}

}

// 只接收channel的函数，chan在右边， <-chan 从chan出来，并接收它
func receiveOnly(ch <-chan int) {
	//ch <- 100 // 不能发送数据到channel, 会报错
	for v := range ch {
		fmt.Println("接收到：", v)
		time.Sleep(500 * time.Millisecond)
	}
}

// 只发送channel的函数，chan在左边， chan<- 发送到chan里面
func sendOnly(ch chan<- int) {
	for i := 0; i < 6; i++ {
		fmt.Println("发送前：", i)
		ch <- i // 缓存满了会阻塞，立马阻塞当前数据，直到有空闲位置
		fmt.Println("发送后：", i)
	}
	i := ch //可以读取 chan, 但是i是一个指针，读取不到任何信息
	fmt.Println("read from sendOnly channel type: ", i)
	//close(ch)
}
