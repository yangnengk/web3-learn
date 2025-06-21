package main

import (
	"fmt"
	"time"
)

// 局部变量
func localVariable(param int) (res int) {
	var decVar int = 10
	fmt.Println("decVar = ", decVar)
	return 20
}

func main() {
	var a int

	fmt.Println("------------------------------if")
	if b := 1; b == 0 {
		fmt.Println("b == 0")
	} else {
		c := 2
		fmt.Println("declare c = ", c)
		fmt.Println("b != 0")
	}

	// fmt.Println(b)	// b和c都未定义
	// fmt.Println(c)

	fmt.Println("------------------------------switch")
	switch d := 3; d {
	case 1:
		e := 4
		fmt.Println("declare e = ", e)
		fmt.Println("d == 1")
	case 2:
		f := 5
		fmt.Println("declare f = ", f)
		fmt.Println("d == 2")
	case 3:
		f := 6
		fmt.Println("declare f = ", f)
		fmt.Println("d == 3")
	default:
		fmt.Println("d != 1 && d != 2 && d != 3")
	}

	// fmt.Println(e)	// e 和 f 都未定义
	// fmt.Println(f)

	fmt.Println("------------------------------for")
	for i := 0; i < 1; i++ {
		forA := 1
		println("forA = ", forA)
	}
	// fmt.Println(forA) // forA 未定义

	fmt.Println("------------------------------select")
	select {
	case <-time.After(time.Second):
		selectA := 1
		println("selectA = ", selectA)
	default:
		selectB := 2
		fmt.Println("selectB = ", selectB)
	}
	// fmt.Println(selectA) // selectA 未定义

	fmt.Println("------------------------------匿名代码块")
	{
		blockA := 1
		fmt.Println("blockA = ", blockA)
	}
	// fmt.Println(blockA) // blockA 未定义

	fmt.Println("------------------------------局部变量")
	fmt.Println("a = ", a)
}
