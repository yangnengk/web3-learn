package main

import (
	"fmt"
)

func main() {
	var p = [][]int{
		{1, 3},
		{8, 10},
		{2, 6},
		{15, 18},
		{2, 8},
	}
	// var p = [][]int{
	// 	{1, 4},
	// 	{2, 3},
	// }
	var res = merge(p)
	fmt.Println("res:", res)
}

func merge(intervals [][]int) [][]int {

	for i := 0; i < len(intervals); i++ {
		var row1 = intervals[i]
		p1 := &intervals[i]
		for j := i + 1; j < len(intervals); j++ {
			var row2 = intervals[j]
			p2 := &intervals[j]
			if (row1[0] > row2[0]) || (row1[0] == row2[0] && row1[len(row1)-1] > row2[len(row2)-1]) {
				*p1, *p2 = *p2, *p1
			}
		}
	}
	// sort.Slice(intervals, func(i, j int) bool {
	// 	return intervals[i][0] < intervals[j][0]
	// })

	var res = [][]int{}
	res = append(res, intervals[0])

	for i := 1; i < len(intervals); i++ {
		slice := res[len(res)-1]
		curr := intervals[i]
		if slice[len(slice)-1] >= curr[0] {
			if slice[len(slice)-1] < curr[len(curr)-1] {
				slice[len(slice)-1] = curr[len(curr)-1]
			}
		} else {
			res = append(res, curr)
		}
	}

	return res
}
