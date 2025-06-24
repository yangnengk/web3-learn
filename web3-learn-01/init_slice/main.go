package main

import "fmt"

/*
切片(Slice)并不是数组或者数组指针，而是数组的一个引用，

切片本身是一个标准库中实现的一个特殊的结构体，这个结构体中有三个属性，分别代表数组指针、长度、容量。
*/
func main() {
	// 方式1，声明与初始化切片，一个空切片
	var s1 []int = []int{}
	fmt.Println("s1:", s1, "len(s1):", len(s1), "cap(s1):", cap(s1))
	// 方式2，声明与初始化切片，类型推导
	var s2 = []int{}
	fmt.Println("s2:", s2, "len(s2):", len(s2), "cap(s2):", cap(s2))
	// 方式3，声明与初始化切片,自动推导类型
	s3 := []int{}
	fmt.Println("s3:", s3, "len(s3):", len(s3), "cap(s3):", cap(s3))
	// 方式4，与方式1、2、3 等价，可以在大括号中定义切片初始元素
	s4 := []int{1, 2, 3, 4}
	fmt.Println("s4:", s4, "len(s4):", len(s4), "cap(s4):", cap(s4))
	// 方式5，用make()函数创建切片，创建[]int类型的切片，指定切片初始长度为0
	s5 := make([]int, 0)
	fmt.Println("s5:", s5, "len(s5):", len(s5), "cap(s5):", cap(s5))
	// 方式6，用make()函数创建切片，创建[]int类型的切片，指定切片初始长度为2，指定容量参数4
	s6 := make([]int, 2, 4)
	fmt.Println("s6:", s6, "len(s6):", len(s6), "cap(s6):", cap(s6))
	// 方式7，引用一个数组，初始化切片
	a := [5]int{1, 2, 3, 4, 5}
	s7 := a[2:]
	fmt.Println("s7:", s7, "len(s7):", len(s7), "cap(s7):", cap(s7))
	// // 从数组下标1开始，直到数组下标3(不包括3)的元素，创建一个新的切片
	s8 := a[1:3]
	fmt.Println("s8:", s8, "len(s8):", len(s8), "cap(s8):", cap(s8))
	// 从0到下标2(不包括2)的元素，创建一个新的切片
	s9 := a[:2]
	fmt.Println("s9:", s9, "len(s9):", len(s9), "cap(s9):", cap(s9))

}
