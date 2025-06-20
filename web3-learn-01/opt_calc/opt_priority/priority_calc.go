package main

import "fmt"

func main() {
	// 优先级
	var a int = 21
	var b int = 10
	var c int = 16
	var d int = 5
	var e int

	e = (a + b) * c / d
	fmt.Println("( a + b ) * c / d 的值为：", e)

	e = ((a + b) * c) / d
	fmt.Println("((a + b) * c) / d 的值为：", e)

	e = (a + b) * (c / d)
	fmt.Println("(a + b) * (c / d) 的值为：", e)

	e = a + (b*c)/d
	fmt.Println("a + (b * c) / d 的值为：", e)

	f := 3 + 4 ^ 3 | 2&2*3<<1
	fmt.Println("f 的值为：", f)
}
