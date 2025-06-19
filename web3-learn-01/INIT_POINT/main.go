package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var p1 *int
	var p2 *string

	i := 1
	s := "Hello"
	// 基础类型数据，必须使用变量名获取指针，无法直接通过字面量获取指针
	// 因为字面量会在编译期被声明为成常量，不能获取到内存中的指针信息
	p1 = &i
	p2 = &s

	p3 := &p2
	fmt.Println("p1指针的值", p1)
	fmt.Println("p2指针的值", p2)
	fmt.Println("p3指针的值", p3)
	fmt.Println("*p3的值", *p3) // *p3的值是s的内存的地址值

	//================================================
	// 除了访问值以外，同样可以通过指针修改原始变量的值。
	// var p1 *int
	// i := 1
	// p1 = &i
	// fmt.Println(*p1 == i)
	// *p1 = 2
	// fmt.Println(i)

	// // 修改指针指向的值
	// fmt.Println("========================================= 修改指针指向的值")
	// a := 2
	// var p *int
	// fmt.Println(&a)
	// p = &a
	// fmt.Println(p, &a)

	// var pp **int
	// pp = &p
	// fmt.Println(pp, p)

	// **pp = 3
	// fmt.Println(pp, *pp, p)
	// fmt.Println(**pp, *p)
	// fmt.Print(&a, a)

	// ------------------------------------指针、unsafe.Pointer 和 uintptr
	// var a int
	// var p *int
	// p = &a
	// p = p + 1 // 无法通过编译
	// fmt.Println(p)

	//  unsafe.Point 类型和 uintptr 类型
	// *T <---> unsafe.Pointer <---> uintptr

	a := "hello world"
	// p := &a
	upA := uintptr(unsafe.Pointer(&a)) // p和&a是一样的效果
	upA += 1
	fmt.Println(upA) // 此处输出一个uintptr类型的值，该值是a的地址加1后的整数值

	c := (*uint8)(unsafe.Pointer(upA)) // 将uintptr类型的值转换为*uint8类型的指针
	fmt.Println(*c)
	fmt.Println(c)
}
