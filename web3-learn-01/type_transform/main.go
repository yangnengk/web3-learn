package main

import (
	"fmt"
	"strconv"
)

// 类型转换
func main() {

	fmt.Println("------------------------数字类型转换")
	var i int32 = 17
	var b byte = 5
	// var f float32

	// 数字类型可以直接转换
	var f = float32(i) / float32(b)
	fmt.Printf("f = %f\n", f)

	// 当int32类型强转成byte时，高位会被直接丢弃

	var i2 int32 = 256
	var b2 byte = byte(i2)
	fmt.Printf("b2 = %d\n", b2)

	// 字符类型转换
	fmt.Println("------------------------字符类型转换")
	str := "hello, 123, 你好"
	var bytes []byte = []byte(str)
	var runes []rune = []rune(str)
	fmt.Printf("bytes = %v\n", bytes)
	fmt.Printf("runes = %v\n", runes)

	str2 := string(bytes)
	str3 := string(runes)
	fmt.Printf("str2 = %v\n", str2)
	fmt.Printf("str3 = %v\n", str3)

	// 但是也会经常有数字与字符串相互转换的需求，这时需要使用到 go 提供的标准库 strconv。
	//strconv 可以把数字转成字符串，也可以把字符串转换成数字
	fmt.Println("------------------------数字与字符串相互转换")
	str4 := "15"
	//num, err := strconv.Atoi(str4)
	//if err != nil {
	//	panic(err) // panic()立刻终止当前函数执行，返回panic的错误信息，直到程序崩溃
	//}
	//fmt.Printf("字符串转成int：num = %d\n", num)
	//
	//str5 := strconv.Itoa(num)
	//fmt.Printf("int转成字符串：str5 = %s\n", str5)

	ui64, err := strconv.ParseUint(str4, 10, 32) // 将10进制字符串转换为无符号64位整数
	if err != nil {
		panic(err)
	}
	i64, err := strconv.ParseInt(str4, 16, 32) // 将16进制字符转成为有符号64位整数
	if err != nil {
		panic(err)
	}
	fmt.Printf("字符串转换成无符号64位整数：ui64 = %d\n", ui64)
	fmt.Printf("字符串转换成有符号64位整数：i64 = %d\n", i64)

	str6 := strconv.FormatUint(ui64, 16) // 将无符号64位整数转成16进制字符串
	str7 := strconv.FormatInt(i64, 16)   // 将有符号64位整数转成16进制字符串
	fmt.Printf("无符号64位整数转成16进制字符串：str6 = %s\n", str6)
	fmt.Printf("有符号64位整数转成16进制字符串：str7 = %s\n", str7)

	str8 := fmt.Sprintf("%d", int32(ui64))
	fmt.Printf("无符号64位整数转成10进制字符串：str8 = %s\n", str8)
	/**
	最常见的转换是字符串与 int 类型之间相互转换。也就是 Atoi 方法与 Itoa 方法。
	当需要把字符串转换成无符号数字时，目前只能转换成 uint64 类型，需要其他位的数字类型需要从 uint64 类型转到所需的数字类型。
	同时可以看到当使用 ParseUint 方法把字符串转换成数字时，或者使用 FormatUint 方法把数字转换成字符串时，都需要提供第二个参数 base，这个参数表示的是数字的进制，即标识字符串输出或输入的数字进制
	*/
}
