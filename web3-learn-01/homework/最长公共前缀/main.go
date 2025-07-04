package main

import "fmt"

func main() {
	// var str = []string{"flower", "flow", "flight"}
	var str = []string{"flower", "flow", "flight"}
	var res = longestCommonPrefix(str)
	fmt.Println("返回结果：", res)
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	} else if len(strs) == 1 {
		return strs[0]
	}
	for i, c := range strs[0] {
		fmt.Println(string(c))
		for j := 1; j < len(strs); j++ {
			if len(strs[j]) == i || byte(c) != strs[j][i] {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}
