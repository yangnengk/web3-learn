package main

import "errors"

func main() {
	//var a int = 10;
	//if b := 1;a > 10 {
	//	b = 2
	//	fmt.Println("a > 10")
	//} else if c := 3;b > 1 {
	//	b = 3
	//	fmt.Println("b > 1")
	//} else {
	//	fmt.Println("其他")
	//	if c == 3 {
	//		fmt.Println("c == 3")
	//	}
	//	fmt.Println(b)
	//	fmt.Println(c)
	//}

	//fmt.Println(b)	// b没有定义
	//fmt.Println(c)	// c没有定义

	if ok := method(); ok {
		// handle normal
	} else if err := methodThrowError(); err != nil {
		// handle error
	}
}

func method() bool {
	return true
}

func methodThrowError() error {
	return errors.New("error msg")
}
