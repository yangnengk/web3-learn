package main

import (
	"fmt"
	"unicode/utf8"

	_ "github.com/learn/init_order/pkg1"
)

const mainName = "main"

var mainVar = getMainVar()

func init() {
	fmt.Println("main init method invoked")
}

func main() {
	// fmt.Println("main method invoked!")

	// 十六进制
	// var a uint8 = 0xF
	// var b uint8 = 0xf

	// // 八进制
	// var c uint8 = 017
	// var d uint8 = 0o17
	// var e uint8 = 0o17

	// // 二进制
	// var f uint8 = 0b1111
	// var g uint8 = 0b1111

	// // 十进制
	// var h uint8 = 15

	// fmt.Println(a == b)
	// fmt.Println(b == c)
	// fmt.Println(c == d)
	// fmt.Println(d == e)
	// fmt.Println(e == f)
	// fmt.Println(f == g)
	// fmt.Println(g == h)

	// float32 是单精度浮点数，float64 是双精度浮点数
	// 对于浮点类型需要被自动推到的变量，其类型都会被自动设置为 float64
	var float1 float32 = 10
	float2 := 10.0
	// 双精度浮点数不能直接赋值给 float32 类型	float1 = float2
	// 正确写法
	float1 = float32(float2)
	fmt.Println(float1 == float32(float2))

	fmt.Println("====================================== complex")
	// 复数：complex64，complext128, 自动推导默认为 complex128
	var c1 complex64
	c1 = 1.10 + 0.1i
	c2 := 1.10 + 0.1i
	c3 := complex(1.10, 0.1)
	fmt.Println(c1 == complex64(c2))
	fmt.Println(complex128(c1) == c2)
	fmt.Println(c2 == c3)
	x := real(c2)
	y := imag(c2)
	z := imag(c3)
	fmt.Println(x, y, z)

	fmt.Println("====================================== byte")
	// byte 是 uint8 的内置别名，可以把 byte 和 uint8 视为同一种类型
	var s string = "hello world"
	var bytes []byte = []byte(s)
	fmt.Println("convert \"hello world\" to []byte", bytes)
	fmt.Println(string(bytes) == s)

	fmt.Println("====================================== rune")
	// rune 是 int32 的内置别名，可以把 rune 和 int32 视为同一种类型。但 rune 是特殊的整数类型，代表一个 Unicode 码点，相当于java中的 char
	var r rune = 'a'
	var r2 rune = '中'
	fmt.Println(r, r2)
	fmt.Println(string(r), string(r2))

	var s1 string = "a b c, 你好世界"
	var runes []rune = []rune(s1)
	fmt.Println(runes)
	fmt.Println(string(runes) == s1)

	fmt.Println("====================================== string")
	// 在 Go 中，字符串是 UTF-8 编码的，并且所有的 Go 源码都必须是 UTF-8 编码。
	//字符串的字面量有两种形式。
	// 一种是解释型字面表示（interpreted string literal，双引号风格），一种是原生型字面表示（raw string literal，反引号风格）。
	var s2 string = "Hello world!"
	var s3 string = `Hello world!`
	fmt.Println(s2 == s3) // 输出结果为 true

	fmt.Println("====================================== byte、rune 与 string 之间的联系")
	var s4 string = "Go语言"
	var bytes1 []byte = []byte(s4)
	var runes1 []rune = []rune(s4)

	fmt.Println("string length:", len(s4))
	fmt.Println("byte length:", len(bytes1))
	fmt.Println("rune length:", len(runes1))
	// 也就是说直接获取字符串的长度是把字符串转换成 []byte 之后 []byte 的长度。
	fmt.Println("direct string length (rune count):", utf8.RuneCountInString(s4)) // 字符串的字符数
	// 进行截取
	fmt.Println("string[0:5]:", s4[0:7])
	fmt.Println("bytes1[0:7]:", string(bytes1[0:7]))
	fmt.Println("runes1[0:3]:", string(runes1[0:3]))

	fmt.Println("====================================== 零值")
	var digit int
	var s5 string // 空字符串
	var b bool
	fmt.Println(digit, s5, b)

}

func getMainVar() string {
	fmt.Println("main.getMainVar method invoked!")
	return mainName
}
