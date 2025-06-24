package main

import "fmt"

// 访问切片
func main() {
	s1 := []int{5, 4, 3, 2, 1}

	// 下标切片访问
	e1 := s1[0]
	e2 := s1[1]
	e3 := s1[2]

	fmt.Println(s1, e1, e2, e3)

	// 向指定位置赋值
	s1[0] = 100
	s1[2] = 1032
	fmt.Println(s1)
	// range迭代访问切片
	for i, v := range s1 {
		fmt.Printf("index and value s1[%d] = %d\n", i, v)
	}

	for i := range s1 {
		fmt.Printf("only index %d\n", i)
	}

	// 切片还可以使用 len() 和 cap() 函数访问切片的长度和容量
	// 当切片是 nil 时，len() 和 cap() 函数获取的到值都是 0
	var nilSlice []int
	fmt.Println("nilSlice:", nilSlice, "len(nilSlice):", len(nilSlice), "cap(nilSlice):", cap(nilSlice))

	slice2 := []int{1, 2, 3, 4, 5}
	fmt.Println("slice2:", slice2, "len(slice2):", len(slice2), "cap(slice2):", cap(slice2))

	//切片添加元素
	// 内置函数 append() 只有切片类型可以使用，第一个参数必须是切片类型，后面追加的元素参数是变长类型，一次可以追加多个元素到切片。并且每次 append() 都会返回一个新的切片引用
	fmt.Println("--------------------------切片添加元素--------------------------")
	sAppend := []int{}
	fmt.Println("sAppend:", sAppend, "len(sAppend):", len(sAppend), "cap(sAppend):", cap(sAppend))
	sAppend = append(sAppend)
	fmt.Println("sAppend:", sAppend, "len(sAppend):", len(sAppend), "cap(sAppend):", cap(sAppend))
	sAppend = append(sAppend, 1)
	fmt.Println("sAppend:", sAppend, "len(sAppend):", len(sAppend), "cap(sAppend):", cap(sAppend))
	sAppend = append(sAppend, 2, 3, 4, 5)
	fmt.Println("sAppend:", sAppend, "len(sAppend):", len(sAppend), "cap(sAppend):", cap(sAppend))
	// 指定位置添加元素
	sAppend2 := []int{1, 2, 4, 5}
	// ...三个点表示将切三拆分为单个元素
	sAppend2 = append(sAppend2[:2], append([]int{3}, sAppend2[2:]...)...)
	fmt.Println("sAppend2:", sAppend2, "len(sAppend2):", len(sAppend2), "cap(sAppend2):", cap(sAppend2))
	// 首位添加元素
	sAppend2 = append([]int{12}, sAppend2...)
	fmt.Println("sAppend2:", sAppend2, "len(sAppend2):", len(sAppend2), "cap(sAppend2):", cap(sAppend2))
	// 移除指定位置元素
	sAppend3 := []int{1, 2, 3, 4, 5}
	sAppend3 = append(sAppend3[0:2], sAppend3[3:]...) // 移除下标为2的元素
	fmt.Println("sAppend3:", sAppend3, "len(sAppend3):", len(sAppend3), "cap(sAppend3):", cap(sAppend3))
	sAppend3 = sAppend3[:2]  // 截取前面部分
	sAppend3 = sAppend3[3:]  // 截取后面部分
	sAppend3 = sAppend3[2:3] // 截取中间部分
}
