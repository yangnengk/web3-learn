package main

import "fmt"

var a int

func main() {
	// {
	// 	fmt.Println("global variable a:", a)
	// 	a = 3
	// 	fmt.Println("global variable a:", a)

	// 	a := 10
	// 	fmt.Println("local variable a:", a)
	// 	a--
	// 	fmt.Println("local variable a:", a)
	// }

	// fmt.Println("global variable a:", a)

	var b int = 4
	fmt.Println("local variable b:", b)
	if b := 3; b == 3 {
		fmt.Println("if statement b:", b)
		b--
		fmt.Println("if statement b:", b)
	}
	fmt.Println("local variable b:", b)
}
