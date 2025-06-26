package main

import (
	"context"
	"fmt"
	"time"
)

// 并发
/*
select 语义是和 channel 绑定在一起使用的，select 可以实现从多个 channel 收发数据，可以使得一个 goroutine 就可以处理多个 channel 的通信。
语法上和 switch 类似，有 case 分支和 default 分支，只不过 select 的每个 case 后面跟的是 channel 的收发操作
1.select 关键字和后面的 { 之间，不能有表达式或者语句。
2.每个 case 关键字后面跟的必须是 channel 的发送或者接收操作
3.允许多个 case 分支使用相同的 channel，case 分支后的语句甚至可以重复
*/
func main() {
	chan1 := make(chan int, 10)
	chan2 := make(chan int, 10)
	chan3 := make(chan int, 10)

	//go func() {
	//	for i := 0; i < 10; i++ {
	//		chan1 <- i
	//		chan2 <- i
	//		chan3 <- i
	//	}
	//}()

	for i := 0; i < 10; i++ {
		chan1 <- i
		chan2 <- i
		chan3 <- i
	}

	// 只要有一个 case 可以执行，select 就会继续执行，否则就一直阻塞
	// 这里只执行10次，随机选取case执行
	//for i := 0; i < 10; i++ {
	//	// chan中读取数据
	//	select {
	//	case x := <-chan1:
	//		fmt.Println("receive from chan1:", x)
	//	//case y := <-chan2:
	//	//	fmt.Println("receive from chan2:", y)
	//	case chan2 <- 11:
	//		// chan容量只有10，上面 go func() 已经把 chan2 塞满，所以这里会阻塞。
	//		// 但是go func是异步的，在go func还未塞满，这里可能会执行， 去掉go func()这里就会阻塞
	//		fmt.Println("send to chan2:", 11)
	//	case z := <-chan3:
	//		fmt.Println("receive from chan3:", z)
	//	}
	//}

	/*
		在执行 select 语句的时候，如果当下那个时间点没有一个 case 满足条件，就会走 default 分支。
		至多只能有一个 default 分支。
		如果没有 default 分支，select 语句就会阻塞，直到某一个 case 满足条件。
		如果 select 里任何 case 和 default 分支都没有，就会一直阻塞。
		如果多个 case 同时满足，select 会随机选一个 case 执行。
	*/
	<-chan1
	<-chan3
	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second) // 2秒后超时

	go func() {
		for {
			// chan中读取数据
			select {
			case chan1 <- 10:
				fmt.Println("send to chan1", 10)
			case chan2 <- 11:
				fmt.Println("send to chan2:", 11)
			case chan3 <- 12:
				fmt.Println("send to chan3:", 12)
			//default:
			//	fmt.Println("no channel ready")
			case <-ctx.Done():
				fmt.Println("context done")
				if err := ctx.Err(); err != nil {
					fmt.Println("context error:", err)
				}
				return
			}
		}
	}()
	fmt.Println("select end")

	time.Sleep(4 * time.Second)
}
