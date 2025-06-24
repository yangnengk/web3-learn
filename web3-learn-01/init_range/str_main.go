package main

import (
	"fmt"
)

/*
这里需要强调一点，在 Go 中，所有字符串都是按照 Unicode 编码的。
第一 for 循环中，遍历了变量 str1，它有六个字符，而且这些字符都可以使用一个 byte 表示，所以循环了六次才退出循环。
而第二个 for 循环中，遍历了变量 str2，它有四个中文字符，那么按照 Unicode 编码的标准，很显然不能仅仅被四个 byte 表示，但它还是只循环了 4 次，恰好是中文字符的长度。
那么也就是说，在 Go 中，遍历字符串时，实际上是在遍历从字符串转换来的 rune 切片，只是恰好在某些时候，字符串转换成 byte 切片和字符串转换成 rune 切片之后的长度相同，
看起来是在逐个遍历 byte 切一样
*/
func main() {
	// 对字符串进行迭代
	str1 := "abc123"
	for index := range str1 {
		fmt.Println("str1 index:", index, "str1 value:", str1[index])
	}

	/*
		当直接使用下标取字符串某个下标位置上的值时，取出来的是 byte 值。
		但是当使用 range 关键字直接获取到某个下标位置的值时，取出的是一个完整的 rune 类型的值。
	*/
	str2 := "测试中文"
	for index, v := range str2 {
		fmt.Println("str2 index:", index, "str2 value:", str2[index], "str2 value:", string(v))
	}
	fmt.Println("str2 len:", len(str2))

	runesFromStr2 := []rune(str2)
	bytesFromStr2 := []byte(str2)

	fmt.Println("runesFromStr2 len:", len(runesFromStr2))
	fmt.Println("bytesFromStr2 len:", len(bytesFromStr2))
}
