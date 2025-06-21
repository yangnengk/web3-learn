package main

import "fmt"

func main() {
	a := A{1}
	// 把方法值赋值给函数变量
	function1 = a.add

	// 声明闭包，并立即执行
	// 此闭包返回值是另一个闭包（带参闭包）
	returnFunc := func() func(int, string) (int, string) {
		fmt.Println("this is anonymous function")
		return func(i int, s string) (int, string) {
			return i, s
		}
	}()

	// 执行returnFunc闭包并传递参数
	ret1, ret2 := returnFunc(1, "hello")
	fmt.Println("call closure function, return1:", ret1, "return2:", ret2)

	fmt.Println("a.i = ", a.i)
	fmt.Println("after call function1, a.i = ", function1(1))
	fmt.Println("a.i = ", a.i)
}

type A struct {
	i int
}

// 定义方法
func (a *A) add(v int) int {
	a.i += v
	return a.i
}

// 声明变量函数
var function1 func(int) int

// 声明闭包
var squart2 func(int) int = func(p int) int {
	p *= p
	return p
}
