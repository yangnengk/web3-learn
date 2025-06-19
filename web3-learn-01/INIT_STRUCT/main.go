package main

import (
	"fmt"
	"sync"
)

func main() {
	var noName = struct {
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
	fmt.Println(noName)
	fmt.Println(noName.int)
	fmt.Println(noName.Name)
	fmt.Println(string(noName.rune))
}

/*
结构体
*/
type Person struct {
	Name  string            `json:"name" gorm:"column:<name>"`
	Age   int               `json:"age" gorm:"column:<name>"`
	Call  func() byte       `json:"-" gorm:"column:<name>"`
	Map   map[string]string `json:"map" gorm:"column:<name>"`
	Ch    chan string       `json:"-" gorm:"column:<name>"`
	Arr   [32]uint8         `json:"arr" gorm:"column:<name>"`
	Slice []interface{}     `json:"slice" gorm:"column:<name>"`
	ptr   *int              `json:"-"`
	once  sync.Once         `json:"-"`
}

type Custom struct {
	field1, field2, field3 byte
}
