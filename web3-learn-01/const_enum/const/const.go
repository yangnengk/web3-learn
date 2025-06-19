package main

type testConst struct {
	a int
	b int
	c int
}

type AliasInt int // 类型别名

// 常量赋值不用带冒号
// 常量定义后不能修改
func main() {
	const a = 10
	const b = 20
	const c = a + b
	const d, e, f = 1, 2, "hello"

	const g bool = true

	const (
		h = 100
		i = 200
		j = "hello 111"
	)

	println(a)
	println(b)
	println(c)
	println(d)
	println(e)
	println(f)
	println(g)
	println(h)
	println(i)
	println(j)

	// 这里会报错，因为常量定义后不能修改，所以结构体不能作为常量赋值
	// const testConst1 = testConst{1, 2, 3}

	const testConst2 = AliasInt(100)
	println(testConst2)

}
