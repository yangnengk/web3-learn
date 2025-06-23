package main

import "fmt"

// 数组访问
func main() {
	a := [5]int{1, 2, 3, 4, 5}

	// 使用下标获取
	element := a[2]
	fmt.Println("element = ", element) //element =  3

	// 使用range遍历
	for i, v := range a {
		fmt.Println("index = ", i, "value = ", v)
		fmt.Printf("a[%d] = %d \n", i, a[i])
	}
	for i := range a {
		fmt.Println("only index = ", i)
	}

	// 读取数组长度
	fmt.Println("len(a) = ", len(a)) //len(a) =  5

	// 使用下标，for循环遍历
	for i := 0; i < len(a); i++ {
		fmt.Println("for i = ", i, ", a[i] = ", a[i])
	}
}
