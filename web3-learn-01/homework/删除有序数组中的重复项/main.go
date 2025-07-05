package main

import "fmt"

func main() {
	var p = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	var res int = removeDuplicates(p)
	fmt.Println("切片大小:", res)
}

func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	var res int = 1
	for l, r := 0, 1; r < len(nums); {
		if nums[l] != nums[r] {
			for i := l + 1; i < r; i++ {
				nums[i] = nums[r]
			}
			l++
			res++
		}
		r++
	}
	fmt.Println("切片:", nums)
	return res
}
