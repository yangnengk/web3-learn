package main

import "fmt"

// 结构体嵌套
type A struct {
	Name string
}

type B struct {
	A
	Age int
}

type C struct {
	B
	Addr string
	Name string
}

func main() {
	a := A{Name: "张三"}
	b := B{A: a, Age: 18}
	c := C{B: b, Addr: "北京", Name: "李四"}

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c.A.Name)
	fmt.Println(c.Name)
	fmt.Println(c.A)
}
