package main

import "fmt"

func main() {
	hash := map[string]int{
		"a": 1,
		"f": 2,
		"z": 3,
		"c": 4,
	}
	// 每次循环遍历，key和value的顺序是随机的
	for key := range hash {
		fmt.Printf("key=%s, value=%d\n", key, hash[key])
	}

	for key, value := range hash {
		fmt.Printf("key=%s, value=%d\n", key, value)
	}
}
