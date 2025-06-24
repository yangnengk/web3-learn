package main

import (
	"fmt"
	"reflect"
)

// 数组遍历
func main() {
	// array := [...]int{1, 2, 3, 4, 5}
	// slice := []int{1, 2, 5}

	// // 方法1：只拿数组的索引
	// for index := range array {
	// 	fmt.Println("array index:", index, "value:", array[index])
	// }

	// for index := range slice {
	// 	fmt.Println("slice index:", index, "value:", slice[index])
	// }

	// // 方法2：拿数组的索引和值
	// for index, value := range array {
	// 	fmt.Println("(index) array index:", index, "value:", array[index])
	// 	fmt.Println("(index and value) array index:", index, "value:", value)
	// }
	// for index, value := range slice {
	// 	fmt.Println("(index) slice index:", index, "value:", slice[index])
	// 	fmt.Println("(index and value) slice index:", index, "value:", value)
	// }

	// 遍历二维数组
	array := [...][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	slice := [][]int{
		{1, 2},
		{3},
	}

	// 方法1：只拿数组的索引
	for index := range array {
		// array[index]是一个一维数组
		fmt.Println(reflect.TypeOf(array[index]))
		fmt.Println("array index:", index, "value:", array[index])
	}
	for index := range slice {
		// slice[index]
		fmt.Println(reflect.TypeOf(slice[index]))
		fmt.Println("slice index:", index, "value:", slice[index])
	}

	// 方法2：拿数组的索引和值
	for index, value := range array {
		fmt.Println("(index and value) array index:", index, "value:", value)
	}

	for index, value := range slice {
		fmt.Println("(index and value) slice index:", index, "value:", value)
	}

	// 双重遍历，拿到每个元素的值
	for row_index, row_value := range array {
		for col_index, col_value := range row_value {
			fmt.Println("array index:", row_index, "col_index:", col_index, "value:", col_value)
		}
	}

	for row_index, row_value := range slice {
		for col_index, col_value := range row_value {
			fmt.Println("slice index:", row_index, "col_index:", col_index, "value:", col_value)
		}
	}
}
