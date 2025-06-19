package anony

/*
定义匿名字段
结构体中的字段不是一定要有字段名，也可以仅定义类型，这种只有类型没有字段名的字段被称为匿名字段。
同一个结构体中相同类型的匿名字段只能同时存在一个，但是可以同时声明多个不同类型的匿名字段。
*/
type Custom struct {
	int
	string
	Other string
}

//  var <Name> = struct {
//     <FiledName1> <type1>
//     <FiledName2> <type2>
//     ...
//     <type3>
//     <type4>
//     ...
// } {}	// {}代表结构体实例化
// 匿名内部类必需立马 {} 实例化掉

/*
匿名内部类完整写法
*/
var Test = struct {
	Name string `<tag1>:"<any string>"`
	Age  int    `<tag2>:"<any string>"`
	int
	rune
}{
	Name: "张三",
	Age:  18,
	int:  1,
	rune: 'A',
}

// 在函数或方法中声明匿名结构体并实例化
// func method() {
//     <var name> := struct {
//         <FieldName1> <type1>
//         <FieldName2> <type2>
//         ...
//         <type3>
//         <type4>
//     } {
//         <FieldName1>: <value1>,
//         <FieldName2>: <value2>,

//         <type3>: <value3>,
//         <type4>: <value4>,
//     }
// }

/*
匿名结构体的主要适用场景：
构建测试数据，单元测试方法中一般会直接声明一个匿名结构体的切片，通过遍历切片测试方法的各个逻辑分支。示例代码可以参考：go-ethereum/core/type/transaction_test.go 的 TestYParityJSONUnmarshalling 方法。
http 处理函数中的 JSON 序列化和反序列化，但是不推荐这么使用，应该定义一个正式的结构体。优点是相比 map[string]interface{}无需检查类型、无需检查 key 是否存在并减少相关的代码检查。
*/
