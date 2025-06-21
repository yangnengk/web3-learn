package main

import "fmt"

func main() {
	// 中断for循环
	for i := 0; i < 5; i++ {
		if i == 3 {
			continue
		}
		fmt.Println("第", i, "次循环")
	}
	// 不适用标记
	fmt.Println("====================== 不使用标记")
	for i := 1; i <= 3; i++ {
		fmt.Printf("不使用标记，外部循环, i = %d\n", i)
		for j := 5; j <= 10; j++ {
			if j == 6 {
				continue
			}
			fmt.Printf("不使用标记，内部循环, j = %d\n", j)
		}
	}
	// 使用标记
	fmt.Println("====================== 使用标记")
outter:
	for i := 1; i <= 3; i++ {
		fmt.Printf("使用标记，外部循环, i = %d\n", i)
		for j := 5; j <= 10; j++ {
			if j == 6 {
				continue outter
			}
			fmt.Printf("使用标记，内部循环, j = %d\n", j)
		}
	}
}
