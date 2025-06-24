package main

import "fmt"

// map集合
/*
在 Go 中，map 集合是无序的键值对集合。相比切片和数组，map 集合对索引的自定义程度更高，可以使用任意类型作为索引，也可以存储任意类型的数据。
但是 map 集合中，存储的键值对的顺序是不确定的。当获取 map 集合中的值时，如果键不存在，则返回类型的零值
*/
func main() {
	var m1 map[string]string
	fmt.Println("m1 len:", len(m1))

	m2 := make(map[string]string)
	fmt.Println("m2 len:", len(m2), "m2:", m2)
	m2["name"] = "张三"
	fmt.Println("m2 len:", len(m2), "m2:", m2)
	v, exist := m2["name"]
	fmt.Println("value:", v, "exist:", exist)

	m3 := make(map[string]string, 10)
	fmt.Println("m3 len:", len(m3), "m3:", m3)

	m4 := map[string]string{} // 声明并初始化，相当于 make(map[string]string, 0)
	fmt.Println("m4 len:", len(m4), "m4:", m4)

	m5 := map[string]string{
		"name": "张三",
		"age":  "18",
		"sex":  "男",
	}
	fmt.Println("m5 len:", len(m5), "m5:", m5)
	m5["address"] = "北京" // 添加元素
	fmt.Println("after modify m5 len:", len(m5), "m5:", m5)

	fmt.Println("---------map集合遍历---------")
	// for range 遍历 map 集合
	for k, v := range m5 {
		fmt.Println("k:", k, "v:", v)
	}

	// 使用内置函数删除 map 集合中的元素
	m5["test"] = "test" // 添加元素
	vName, vExist := m5["test"]
	fmt.Println("before delete m5 len:", len(m5), "m5:", m5, "vName:", vName, "vExist:", vExist)
	delete(m5, "test")
	vName, vExist = m5["test"]
	fmt.Println("after delete m5 len:", len(m5), "m5:", m5, "vName:", vName, "vExist:", vExist)

	// 在遍历map集合时，删除某个元素,使用的iterate map,删除某个元素
	for k, v := range m5 {
		fmt.Println("iterate map, will delete key:", k, "value:", v)
		fmt.Println("after delete m5 len:", len(m5), "m5:", m5)
		delete(m5, k)
	}
	fmt.Println("m5:", m5)
}
