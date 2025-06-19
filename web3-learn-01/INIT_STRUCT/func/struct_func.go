package main

import "fmt"

type A struct {
	Name string
	Age  int
}

func (a A) getName() string {
	return a.Name
}

// 这个方法只是A结构体的副本传到了该方法中，所以修改无效
func (a A) setName(name string) {
	a.Name = name
}

// 这个方法修改的是结构体的指针，所以修改有效
func (a *A) setPName(name string) {
	a.Name = name
}

func (a A) getAge() int {
	return a.Age
}

func valueName(a A, name string) {
	a.Name = name
}

func pointerName(a *A, name string) {
	a.Name = name
}

func main() {
	a := A{Name: "张三", Age: 18}

	pa := &a

	// go 里面直接结构体修改是不行的
	a.setName("aaaa")
	fmt.Println(a.getName())
	// go 里面修改结构体的值要用指针
	a.setPName("bbbb")
	fmt.Println(a.getName())

	// 直接通指针修改
	pa.Name = "cccc"
	fmt.Println(a.getName())

	// 此处调用函数直接传结构体，修改无效
	valueName(a, "anyone")
	fmt.Println(a.getName())
	// 此处调用函数传指针，修改有效
	pointerName(&a, "anyone111111")
	fmt.Println(a.getName())

	fmt.Println(a.getAge())
}
