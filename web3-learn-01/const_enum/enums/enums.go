package main

import "fmt"

// type Gender string // string的别名
type Gender byte // string的别名

// const (
// 	Male   Gender = "男"
// 	Female Gender = "女"
// )

const (
	Male Gender = iota
	Female
)

func (g *Gender) isFemale() bool {
	return *g == Female
}

func (g *Gender) isMale() bool {
	return *g == Male
}

func main() {
	// 枚举
	// 枚举类型是预定义的类型，通常也被称为枚举

	var gender = Male

	fmt.Println(gender)
	fmt.Println(gender.isFemale())
	fmt.Println(gender.isMale())
	fmt.Println("-------------------------------------------------------- iota 关键字")
	fmt.Println(January)
	fmt.Println(February)
	fmt.Println(March)
	fmt.Println(April)

	println("-------------------------------------------------")
	println(StateNew)
	println(StateActive)
	println(StateClosed)

}

/*
除了上面的别名类型来声明枚举类型以外，还可以使用 iota 关键字，来自动为常量赋值。
所有常量都自动赋值，从 0 开始，依次加 1。
iota 是行号的计数器，从 0 开始，每遇到一行就加 1。
*/
const (
	January = iota + 1
	February
	March
	April
	// 以此类推
	May
	June
	July
	August
	September
	October
	November
	December
)

type ConnState int

// 如果 iota 定义在 const 定义组中的第 n 行，那么 iota 的值为 n - 1。所以一定要注意 iota 出现在定义组中的第几行，而不是当前代码中它第几次出现
const (
	StateNew    ConnState = -1
	StateNew2   ConnState = -1
	StateActive           = iota
	StateIdle
	StateHijacked
	StateClosed
)
