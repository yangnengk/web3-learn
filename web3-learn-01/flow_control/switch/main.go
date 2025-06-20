package main

import (
	"fmt"
)

/*
@Time : 2021/11/23 17:05
@Description: switch 语句
*/
func main() {
	a := "test string"
	// 1. 基本用法
	switch a {
	case "test":
		fmt.Println("a = ", a)
	case "s":
		fmt.Println("a = ", a)
	case "t", "test string":
		fmt.Println("catch in a test, a = ", a)
	case "n":
		fmt.Println("a = not")
	default:
		fmt.Println("default case ")
	}

	// 变量b仅在当前switch代码块内有效
	switch b := 5; b {
	case 1:
		fmt.Println("b = 1")
	case 2:
		fmt.Println("b = 2")
	case 3, 4:
		fmt.Println("b = 3 or 4")
	case 5:
		fmt.Println("b = 5")
	default:
		fmt.Println("b = ", b)
	}

	// 不指定判断变量，直接在case 后面写条件
	a = "t1"
	b := 5
	switch {
	case a == "t":
		fmt.Println("a = t")
	case b == 3:
		fmt.Println("b = 3")
	case b == 5, a == "test string":
		fmt.Println("b = 5 or a = test string")
	default:
		fmt.Println("default case")
	}

	var d interface{}
	e := CustomType{}
	d = &e
	switch t := d.(type) {
	case byte:
		fmt.Println("d is byte type, ", t)
	case *byte:
		fmt.Println("d is byte point type, ", t)
	case *int:
		fmt.Println("d is int point type, ", t)
	case *string:
		fmt.Println("d is string point type, ", t)
	case *CustomType:
		fmt.Println("d is CustomType point type, ", t)
	case CustomType:
		fmt.Println("d is CustomType type, ", t)
	default:
		fmt.Println("d is unknown type, ", t)

	}
}

type CustomType struct{}
