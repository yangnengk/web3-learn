package main

import "fmt"

func main() {
	var param = []int{2, 7, 11, 15}
	var targe = 9
	var res = twoSum(param, targe)
	fmt.Println("res=", res)
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for i, v := range nums {
		diffNum := target - v
		if diffIndex, ok := m[diffNum]; ok {
			return []int{diffIndex, i}
		}
		m[nums[i]] = i
	}
	return nil
}
