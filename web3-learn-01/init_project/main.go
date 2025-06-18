package main

import (
	"fmt"
	"os"
)

func test() {
	fmt.Print("1111")
}

func main() {
	fmt.Println("hello world")
	fmt.Println(os.Args)
	test()
}
