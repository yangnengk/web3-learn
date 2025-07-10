package main

import "fmt"

func main() {
	/*
		题目 ：编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值
	*/
	var num int64 = 12
	fmt.Println("num的值增加10前：", num)
	intPoint(&num)
	fmt.Println("num的值增加10后：", num)

	/*
		题目 ：实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
		考察点 ：指针运算、切片操作
	*/
	var numSlice = []int64{1, 2, 3, 4, 5}
	fmt.Println("numSlice的值修改前：", numSlice)
	intSlicePoint(&numSlice)
	fmt.Println("numSlice的值修改后：", numSlice)

}

func intPoint(num *int64) {
	*num += 10
}

func intSlicePoint(numSlice *[]int64) {
	for i, _ := range *numSlice {
		(*numSlice)[i] *= 2
	}
}
