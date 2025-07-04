package main

import "fmt"

func main() {
	var p = []int{9}
	var res = plusOne(p)
	fmt.Println("返回结果：", res)
}

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i] += 1
		if digits[i] == 10 {
			digits[i] = 0
		} else {
			return digits
		}
	}
	digits = make([]int, len(digits)+1)
	digits[0]++
	return digits
}
