package main

import (
	"fmt"
	"strconv"
)

/*
接口类型转换
*/
func main() {
	var i interface{} = 3
	//a, ok := i.(int32) // 会报错
	a, ok := i.(int)
	fmt.Printf("%d is a int\n", a)
	if ok {
		fmt.Println("a = ", a)
	} else {
		fmt.Println("convert failed")
	}

	// switch 语句判断结构体接口类型
	var i2 interface{} = byte(1)
	switch t := i2.(type) {
	//.(type) 用于判断接口类型,只能在 switch 语句中使用
	case int:
		fmt.Println("i2 is a int", t)
	case string:
		fmt.Println("i2 is a string", t)
	default:
		fmt.Println("i2 is a unknown type", t)
	}

	// 结构体类型转换
	fmt.Println("------------------------结构体类型转换")
	var s Supplier = &DigitSupplier{value: 1}
	fmt.Println(s.Get())

	// 将接口类型转换成结构体类型（把一个接口类型转换成具体的结构体接口类型）
	b, ok := s.(*DigitSupplier)
	fmt.Println("b = ", b, "ok = ", ok)
}

type Supplier interface {
	Get() string
}

type DigitSupplier struct {
	value int
}

func (i *DigitSupplier) Get() string {
	//return fmt.Sprintf("%d", i.value) // 数字转字符串
	return strconv.FormatInt(int64(i.value), 10) // 数字转字符串
}
