package main

import (
	"fmt"
)

func main() {
	// 数学运算
	// a, b := 1, 2
	// sum := a + b
	// sub := a - b
	// mul := a * b
	// div := a / b
	// mod := a % b

	// fmt.Println(sum, sub, mul, div, mod)
	// fmt.Println("------------------------------------")

	/*
		a := 1
		// 正确写法
		a++
		a--

		// 错误的使用方式
		++a
		--a

		// 错误使用方式，不可以自增时计算,也不能赋值
		b := a++ + 1
		c := a--
	*/

	// 当不同的数字类型混合计算时，必须先把它们转换成同一类型才可以计算：
	// a := 10 + 0.1	// 编译不通过
	// b := byte(1) + 1	// 编译通过
	// sum := a + float64(b) // 编译通过
	// sub := byte(a) - b	// 编译不通过
	// mul := a * float64(b)	// 编译不通过
	// div := int(a) / b		// 编译不通过

	// 关系运算符
	// a := 1
	// b := 5
	// fmt.Println(a == b)
	// fmt.Println(a != b)
	// fmt.Println(a > b)
	// fmt.Println(a < b)
	// fmt.Println(a >= b)
	// fmt.Println(a <= b)

	// 逻辑运算符
	// a := true
	// b := false
	// fmt.Println(a && b)
	// fmt.Println(a || b)
	// fmt.Println(!(a && b))

	// // 位运算符
	// fmt.Println("--------------------------位运算")
	// fmt.Println(0 & 0)
	// fmt.Println(0 | 0)
	// fmt.Println(0 ^ 0)

	// fmt.Println(0 & 1)
	// fmt.Println(0 | 1)
	// fmt.Println(0 ^ 1)

	// fmt.Println(1 & 1)
	// fmt.Println(1 | 1)
	// fmt.Println(1 ^ 1)

	// fmt.Println(1 & 0)
	// fmt.Println(1 | 0)
	// fmt.Println(1 ^ 0)

	// 赋值运算符
	a, b := 1, 2
	var c int
	c = a + b
	fmt.Println("c = a +b, c=", c)

	plusAssignment(c, a)
	subAssignment(c, a)
	mulAssignment(c, a)
	divAssignment(c, a)
	modAssignment(c, a)
	leftMoveAssignment(c, a)
	rightMoveAssignment(c, a)
	andAssignment(c, a)
	orAssignment(c, a)
	norAssignment(c, a)

	// 其他运算符
	a = 4
	var ptr *int
	fmt.Println(a)
	ptr = &a
	fmt.Println("ptr = &a, ptr=", *ptr)

}

func plusAssignment(c, a int) {
	c += a // c = c + a
	fmt.Println("c += a, c=", c)
}

func subAssignment(c, a int) {
	c -= a // c = c - a
	fmt.Println("c -= a, c=", c)
}

func mulAssignment(c, a int) {
	c *= a // c = c * a
	fmt.Println("c *= a, c=", c)
}

func divAssignment(c, a int) {
	c /= a // c = c / a
	fmt.Println("c /= a, c=", c)
}

func modAssignment(c, a int) {
	c %= a // c = c % a
	fmt.Println("c %= a, c=", c)
}

func leftMoveAssignment(c, a int) {
	c <<= a // c = c << a
	fmt.Println("c <<= a, c=", c)
}

func rightMoveAssignment(c, a int) {
	c >>= a // c = c >> a
	fmt.Println("c >>= a, c=", c)
}

func andAssignment(c, a int) {
	c &= a // c = c & a
	fmt.Println("c &= a, c=", c)
}

func orAssignment(c, a int) {
	c |= a // c = c | a
	fmt.Println("c |= a, c=", c)
}

func norAssignment(c, a int) {
	c ^= a // c = c ^ a
	fmt.Println("c ^= a, c=", c)
}
