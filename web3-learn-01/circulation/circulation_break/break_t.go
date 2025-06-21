package main

import (
	"fmt"
	"time"
)

// break 语句
func main() {
	// 中断循环
	for i := 0; i < 5; i++ {
		if i == 3 {
			break
		}
		fmt.Println("第", i, "次循环")
	}

	fmt.Println("====================== switch break")
	// 中断 switch
	switch i := 1; i {
	case 1:
		fmt.Println("进入 case 1")
		if i == 1 {
			break
		}
		fmt.Println("i等于1")
	case 2:
		fmt.Println("进入 case 2")
	default:
		fmt.Println("进入 default")
	}
	fmt.Println("====================== select break")
	// 中断 select
	select {
	case <-time.After(time.Second * 2):
		fmt.Println("过了2秒")
	case <-time.After(time.Second):
		fmt.Println("过了1秒")
		if true {
			break
		}
		fmt.Println("break 之后")

	}

	// 不使用标记
	fmt.Println("====================== 不使用标记")
	for i := 1; i <= 3; i++ {
		fmt.Printf("不使用标记，外部循环, i = %d\n", i)
		for j := 5; j <= 10; j++ {
			fmt.Printf("不使用标记，内部循环, j = %d\n", j)
			if j == 6 {
				break
			}
		}
	}

	// 使用标记
	fmt.Println("====================== 使用标记")
outter:
	for i := 1; i <= 3; i++ {
		fmt.Println("使用标记，外部循环, i = ", i)
		for j := 5; j <= 10; j++ {
			fmt.Println("使用标记，内部循环, j = ", j)
			if j == 6 {
				break outter // 跳出 outter 标签对应的循环
			}
		}
	}
}
