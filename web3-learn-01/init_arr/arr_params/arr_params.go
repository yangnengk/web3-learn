package main

import "fmt"

/*
数组的部分特性类似基础数据类型，当数组作为参数传递时，在函数中并不能改变外部实参的值。
如果想要修改外部实参的值，需要把数组的指针作为参数传递给函数, 跟结构体对象一样
总结：用指针在go语言里，操作的就是实参本身，跟地址什么的没啥关系
*/

type Custom struct {
	a int
}

func main() {

	// a := [5]int{1, 2, 3, 4, 5}
	// fmt.Println("a = ", a)

	// receiveArray(a)
	// fmt.Println("after receiveArray, a = ", a)

	// receiveArrayPointer(&a)
	// fmt.Println("after receiveArrayPointer, a = ", a)

	// i1 := 100
	// i2 := 200
	// i3 := 300
	// i4 := 400
	// i5 := 500

	i1 := Custom{a: 100}
	i2 := Custom{a: 200}
	i3 := Custom{a: 300}
	i4 := Custom{400}
	i5 := Custom{500}

	// b := [5]*int{&i1, &i2, &i3, &i4, &i5}
	// b := [5]*Custom{&i1, &i2, &i3, &i4, &i5}
	// b := [5]Custom{i1, i2, i3, i4, i5}
	b := [5]*Custom{&i1, &i2, &i3, &i4, &i5}
	receiveArrayPointer4(b)
	// fmt.Println("arr point element, b = ", b)
	for i := range b {
		// fmt.Println("arr point element, *b[", i, "] = ", *b[i])
		fmt.Printf("in main func, params[%d] = %p, value = %v \n", i, &b[i], *b[i])
	}
}

func receiveArray(params [5]int) {
	fmt.Println("in receiveArray func, before modify, params = ", params)
	params[1] = 100
	fmt.Println("in receiveArray func, after modify, params = ", params)
}

func receiveArray2(params [5]Custom) {
	fmt.Println("in receiveArray2 func, before modify, params = ", params)
	params[1].a = -111
	fmt.Println("in receiveArray2 func, after modify, params = ", params)
}

// 传入数组指针
func receiveArrayPointer(params *[5]int) {
	fmt.Println("in receiveArrayPointer func, before modify, params = ", params)
	params[1] = 100
	fmt.Println("in receiveArrayPointer func, after modify, params = ", params)
}

func receiveArrayPointer2(params *[5]Custom) {
	fmt.Println("in receiveArrayPointer2 func, before modify, params = ", params)
	params[1].a = -111
	fmt.Println("in receiveArrayPointer2 func, after modify, params = ", params)
}

func receiveArrayPointer3(params *[5]*Custom) {
	params[1].a = -111
	for i, _ := range params {
		fmt.Printf("in receiveArrayPointer3 func, params[%d] = %p, value = %v \n", i, &params[i], *params[i])
	}
	// fmt.Println("in receiveArrayPointer3 func, before modify, params = ", params)
	// params[1].a = -111
	// fmt.Println("in receiveArrayPointer3 func, after modify, params = ", params)
}

func receiveArrayPointer4(params [5]*Custom) {
	params[1].a = -111
	for i, _ := range params {
		fmt.Printf("in receiveArrayPointer4 func, params[%d] = %p, value = %v \n", i, &params[i], *params[i])
	}
	// fmt.Println("in receiveArrayPointer3 func, before modify, params = ", params)
	// params[1].a = -111
	// fmt.Println("in receiveArrayPointer3 func, after modify, params = ", params)
}
