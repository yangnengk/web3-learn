package main

import "fmt"

// 多维数组
func main() {
	// 二维数组
	a := [3][2]int{
		{1, 2},
		{3, 4},
		{5, 6},
	}
	fmt.Println("a = ", a)

	// 三维数组
	b := [3][2][2]int{
		{
			{1, 2},
			{3, 4},
		},
		{
			{5, 6},
			{7, 8},
		},
		{
			{9, 10},
			{11, 12},
		},
	}
	fmt.Println("b = ", b)

	// 可以省去各个位置的初始化，在后续代码中赋值
	var c [3][3][3]int
	d := [3][3][3]int{}

	fmt.Println("c = ", c)
	fmt.Println("d = ", d)

	c[2][1][1] = 100
	c[2][2][1] = 200
	fmt.Println("c = ", c)

	// 多维数组的访问
	fmt.Println("-------------------多维数组的访问-------------------")

	e := [3][2][2]int{
		{
			{1, 2},
			{3, 4},
		},
		{
			{5, 6},
			{7, 8},
		},
		{
			{9, 10},
			{11, 12},
		},
	}

	layer1 := e[0]
	layer2 := e[1]
	layer3 := e[2]
	element := e[0][1][1]
	fmt.Println("layer1 = ", layer1)
	fmt.Println("layer2 = ", layer2)
	fmt.Println("layer3 = ", layer3)
	fmt.Println("element = ", element)

	fmt.Println("-------------------多维数组的遍历-------------------")
	// 多维数组的遍历, 需要嵌套for循环
	for i, v := range e {
		fmt.Println("index = ", i, "value = ", v)
		for j, inner := range v {
			fmt.Println("inner, index = ", j, "value = ", inner)
			for k, inner2 := range inner {
				fmt.Println("inner2, index = ", k, "value = ", inner2)
			}
		}
	}

}
