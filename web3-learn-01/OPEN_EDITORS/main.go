package main

import "fmt"

// var s1 string = "hello world"
// var zero int
// var b1 = true

// var m map[string]int
// var arr [2]byte
// var slice []int
// var p *int

// var (
// 	i  int = 123
// 	b2 bool
// 	s2 = "test"
// )

// var (
// 	group1      = 2
// 	group2 byte = 1
// )

func main() {
	method()
	// fmt.Println(s1)
	// fmt.Println(zero)
	// fmt.Println(b1)
	// fmt.Println(i)
	// fmt.Println(b2)
	// fmt.Println(s2)
	// fmt.Println(group1)
	// fmt.Println(group2)
	// fmt.Println(m)
	// fmt.Println(arr)
	// fmt.Println(p)
	// m = make(map[string]int, 0) // 0表示初始容量
	// m["1"] = 1
	// fmt.Println(m)
	// slice = append(slice, 1)
	// fmt.Println(slice)
	// method1()
	// method2()
	// method3()
	// method4()
}

// func method1() {
// 	// 方式1 类型推导
// 	a := 2
// 	// 方式2 完整的变量声明写法
// 	var b int = 4
// 	// 3方式3，仅声明变量，但是不赋值
// 	var c int
// 	fmt.Println(a, b, c)
// }

// // 方式4，直接在返回值中声明
// func method2() (a int, b string) {
// 	// 这种方式必须声明return关键字
// 	// 并且同样不需要使用，并且也不用必须给这种变量赋值
// 	return 1, "test"
// }

// func method3() (a int, b string) {
// 	a = 1
// 	b = "test"
// 	return
// }

// func method4() (a int, b string) {
// 	return
// }

var a, b, c int = 1, 2, 3

var e, f, g int

var h, i, j = 1, 2, "test"

func method() {
	var k, l, m int = 1, 2, 3
	var n, o, p int
	q, r, s := 1, 2, "test"
	fmt.Println(k, l, m, n, o, p, q, r, s)
}
