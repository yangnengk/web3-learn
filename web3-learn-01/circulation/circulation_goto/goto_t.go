package main

import "fmt"

// 注意代码 label 声明之后，在代码中必须使用到，否则编译时会提示 label xxx defined and not used
func main() {

	gotoPreset := false

preset:
	a := 5

process:
	if a > 0 {
		a--
		fmt.Println("当前a值为：", a)
		goto process
	} else {
		if !gotoPreset {
			gotoPreset = true
			goto preset
		} else {
			goto post
		}
	}

post:
	fmt.Println("main函数结束, 当前a值为：", a)

}
