package main

import "fmt"

func main() {
	var nums = []int{2, 2, 1}
	var res = singleNumber(nums)
	fmt.Println(res)
}

func singleNumber(nums []int) int {
	var m = make(map[int]int, 0)
	for _, v := range nums {
		if m[v] != 0 {
			m[v] += 1
		} else {
			m[v] = 1
		}
	}
	var res int
	for k, v := range m {
		if v == 1 {
			res = k
			break
		}
	}
	return res
}
