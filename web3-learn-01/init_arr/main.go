package main

import "fmt"

func main() {
	// 仅声明
	var a [5]int
	fmt.Println("a = ", a)

	var maArr [2]map[string]string
	fmt.Println("maArr = ", maArr)
	// map的零值是nil，虽然打印出来是非空值，但真实的值是nil
	// maArr[0]["test"] = "1"

	// 声明以及初始化
	var b [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println("b = ", b)

	// 类型推导声明方式
	var c = [5]int{1, 2, 3, 4, 5}
	fmt.Println("c = ", c)

	d := [4]int{1, 2, 3, 4}
	fmt.Println("d = ", d)
	// 使用 ... 代替数组长度,自动计算数组长度
	autoLen := [...]string{"auto1", "auto2", "auto3", "auto4"}
	fmt.Println("autoLen = ", autoLen)
	// print := func(saArr [3]string) { // 会报错：cannot use autoLen (variable of type [4]string) as [3]string value in argument to print
	// 	fmt.Println("saArr = ", saArr)
	// }
	// print(autoLen)

	// 声明时初始化并制定下标元素
	// positionInit := [5]string{1: "position1", 3: "position3"}
	positionInit := [5]int{1: 5, 3: 34}
	fmt.Println("positionInit = ", positionInit) //positionInit =  [0 5 0 34 0]

	// 初始化时，元素个数不能超过生命长度
	// overLen := [5]int{1, 2, 3, 4, 5, 6} // 报错 index 5 is out of bounds (>= 5)
	// fmt.Println("overLen = ", overLen)  //overLen =  [1 2 3 4 5]
}
