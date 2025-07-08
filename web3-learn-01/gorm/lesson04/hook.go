package lesson04

import (
	"fmt"
	"gorm.io/gorm"
)

type Parent struct {
	ID   int `gorm:"primaryKey"`
	Name string
}

type Child struct {
	ID   int `gorm:"primaryKey"`
	Name string
	Parent
}

func (parent *Parent) BeforeCreate(tx *gorm.DB) (err error) {
	fmt.Println("Parent BeforeCreate")
	return nil
}

//func (child *Child) BeforeCreate(tx *gorm.DB) (err error) {
//	fmt.Println("Child BeforeCreate")
//	return nil
//}

func RunHook(db *gorm.DB) {
	db.AutoMigrate(&Parent{}, &Child{})

	// 子类如果有钩子则会覆盖父类的钩子
	db.Create(&Child{Name: "Tom", Parent: Parent{Name: "Jerry"}})
}
