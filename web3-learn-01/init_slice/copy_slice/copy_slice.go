package main

import "fmt"

func main() {
	// 可以使用内置函数 copy() 把某个切片中的所有元素复制到另一个切片，复制的长度是它们中最短的切片长度
	// 但是源切片的长度和容量必须小于等于目标切片的长度和容量,  不然数据会丢失
	src1 := []int{1, 2, 3}
	dist1 := make([]int, 4, 5)

	src2 := []int{1, 2, 3, 4, 5}
	dist2 := make([]int, 3, 3)

	fmt.Println("befor copy src1:", src1, "dist1:", dist1)
	fmt.Println("befor copy src2:", src2, "dist2:", dist2)

	copy(dist1, src1)
	copy(dist2, src2)

	src1[0] = 100 // 修改源切片，不会影响目标切片

	fmt.Println("after copy src1:", src1, "dist1:", dist1)
	fmt.Println("after copy src2:", src2, "dist2:", dist2)
}
