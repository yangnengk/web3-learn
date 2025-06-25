package main

import "fmt"

// 结构体类型转换
/*
结构体类型之间在一定条件下也可以转换的。
1.字段顺序完全相同
2.当两个结构体中的字段名称相同
3.类型都完全相同，
仅结构体名称不同时，这两个结构体类型即可相互转换
*/
func main() {
	a := SimpleFieldA{
		name:  "a",
		value: 1,
	}

	b := SimpleFieldB(a)
	fmt.Printf("convert SimpleFieldA to SimpleFieldB, value is %d \n", b.getValue())

	//只能结构体类型实例之间转换，指针不可以相互转换
	var c interface{} = &a
	_, ok := c.(*SimpleFieldB)
	fmt.Println("c is *SimpleFieldB:", ok)
}

type SimpleFieldA struct {
	name  string
	value int
}

type SimpleFieldB struct {
	name  string
	value int
}

func (s *SimpleFieldB) getValue() int {
	return s.value
}
