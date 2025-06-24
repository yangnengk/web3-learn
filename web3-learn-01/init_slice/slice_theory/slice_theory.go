package main

import "fmt"

// 切片底层原理
// 切片类型实际上是比较特殊的指针类型，当声明一个切片类型时，就是声明了一个指针
// 这个指针指向的切片结构体，切片结构体中记录的三个属性：数组指针、长度、容量。这几个属性在创建一个切片时就定义好，并且在之后都不能再被修改
func main() {
	// s := make([]int, 3, 6)
	// fmt.Println("s:", s, "len:", len(s), "cap:", cap(s))
	// s[1] = 2
	// fmt.Println("set position 1 to 2, s = ", s)

	// // modifySlice(s) // 经过modifySlice函数后，s的值被修改, 说明s是个特殊的指针
	// // fmt.Println("after modifySlice, s = ", s)

	// // 在不使用 append() 函数的情况下，在函数内部对切片的修改，都会影响到原始实例
	// // 使用 append()函数时，需要分两种情况：
	// // 当没有触发切片扩容时:
	// /*
	// 	当使用 append() 函数之后。
	// 	原来的切片引用，长度和容量不变，新追加的值超过切片可访问范围，访问不到新追加的值。
	// 	新的切片引用，与原始切片引用相比，长度加一，容量不变，可以访问到新追加的值。
	// 	在方法内，使用原始切片作为参数，使用 append() 函数追加元素后，同样会创建一个新的切片引用，新追加的值会覆盖之前的值。
	// 	并且修改这个切片，其修改同样会反应到原始切片以及新的切片引用上。
	// */
	// s2 := append(s, 4)
	// fmt.Println("after append, s2 len:", len(s2), "cap:", cap(s2), "s2 = ", s2)
	// fmt.Println("after append, s = ", s)

	// s[0] = 1024
	// fmt.Println("after set s position 0 to 1024, s = ", s)
	// fmt.Println("after set s position 0 to 1024, s2 = ", s2)

	// appendInfunc(s)
	// fmt.Println("after append s in func, s = ", s)
	// fmt.Println("after append s in func, s2 = ", s2)

	// --------------------------------------
	// 当触发了切片扩容时:
	/*
		当 append() 函数触发扩容后，实际上是新创建了一个数组实例，把原来的数组中的数据复制到了新数组中，然后创建一个新的切片实例并返回。
		这时原始切片中持有的数组指针指向的数组与新切片引用中的数组指针指向的数组是两个不同的数组，修改并不会相互影响。
		切片触发扩容前，切片一直共用相同的数组；
		切片触发扩容后，会创建新的数组，并复制这些数据；
		切片本身是一个特殊的指针，go 针对切片类型添加了一些语法糖，方便使用.
	*/
	s := make([]int, 2, 2)
	fmt.Println("initial s = ", s)

	s2 := append(s, 4)
	fmt.Println("after append s2 len:", len(s2), "cap:", cap(s2), "s2 = ", s2)
	fmt.Println("after append s, len:", len(s), "cap:", cap(s), "s = ", s)

	s[0] = 1024
	fmt.Println("after set s position 0 to 1024, s = ", s)
	fmt.Println("after set s position 0 to 1024, s2 = ", s2)

	appendInfunc2(s2)
	fmt.Println("after append s2 in func, s2 = ", s2)

}

func modifySlice(params []int) {
	params[0] = 1024
}

func appendInfunc(param []int) {
	param = append(param, 1022)
	fmt.Println("after append in func, param = ", param)
	param[2] = 112
	fmt.Println("set position 2 to 112, param = ", param)

}

func appendInfunc2(param []int) {
	param1 := append(param, 1022)
	param2 := append(param1, 1025)
	fmt.Println("after append in func, param1 = ", param1)
	param2[2] = 1500
	fmt.Println("set position 2 to 1500, param2 = ", param2)

}
