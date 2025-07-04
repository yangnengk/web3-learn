package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x = 123
	var b = isPalindrome(x)
	if b {
		println("是回文数")
	} else {
		println("不是回文数")
	}
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	xs := strconv.Itoa(x)
	fmt.Println(xs)
	// var r = []rune(xs)
	for i := 0; i < len(xs)/2; i++ {
		one, ok1 := strconv.Atoi(string(xs[i]))
		two, ok2 := strconv.Atoi(string(xs[len(xs)-i-1]))
		fmt.Println(ok1, ok2)
		if ok1 != nil || ok2 != nil {
			return false
		}
		// if xs[i] != xs[len(xs)-i-1] {
		// 	return false
		// }
		if one != two {
			return false
		}
	}

	return true
}
