package main

import (
	"fmt"
)

/*
for循环
*/
func main() {

	//方式一
	for i := 0; i < 10; i++ {
		fmt.Println("方式一， 第", i+1, "次循环")
	}
	fmt.Println("===========================================")
	// 方式
	b := 1
	for b < 10 {
		fmt.Println("方式二， 第", b, "次循环")
		b++
	}

	// 方式三 无限循环
	// ctx, _ := context.WithDeadline(context.Background(), time.Now().Add(time.Second*2)) // 创建一个2秒后超时的上下文
	// var started bool
	// var stopped atomic.Bool
	// for {
	// 	if !started {
	// 		started = true
	// 		go func() {
	// 			for {
	// 				select {
	// 				case <-ctx.Done():
	// 					stopped.Store(true)
	// 					return
	// 				}
	// 			}
	// 		}()
	// 	}
	// 	fmt.Println("main")
	// 	if stopped.Load() {
	// 		break
	// 	}
	// }

	// 遍历数组
	fmt.Println("------------------------------------------遍历数组")
	var a [10]string
	a[0] = "hello world"
	for i := range a {
		fmt.Println("当前下标为：", i, "，值为：", a[i])
	}
	for i, e := range a {
		fmt.Println("当前下标为：", i, "，值为：", e)
	}

	// 遍历切片
	fmt.Println("------------------------------------------遍历切片")
	s := make([]string, 10)
	s[0] = "Hello"
	for i := range s {
		fmt.Println("当前下标为：", i, "，值为：", s[i])
	}
	for i, e := range s {
		fmt.Println("当前下标为：", i, "，值为：", e)
	}

	// 遍历map
	fmt.Println("------------------------------------------遍历map")
	m := make(map[string]string, 10)
	m["a"] = "Hello a"
	m["b"] = "Hello b"
	m["c"] = "Hello c"
	for i := range m {
		fmt.Println("当前key：", i, "，值为：", m[i])
	}
	for i, e := range m {
		fmt.Println("当前key：", i, "，值为：", e)
	}
}
