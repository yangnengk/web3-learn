package lesson04

import (
	"fmt"
	"gorm.io/gorm"
)

type Dog struct {
	ID   int
	Name string
	Toy  Toy `gorm:"polymorphic:Owner"`
}

type Cat struct {
	ID   int
	Name string
	Toy  Toy `gorm:"polymorphic:Owner"`
}

type Toy struct {
	ID        int
	Name      string
	OwnerType string
	OwnerID   int
}

func Run(db *gorm.DB) {
	db.AutoMigrate(&Dog{}, &Cat{}, &Toy{})

	// 多态
	//db.Create(&Dog{Name: "Tom", Toy: Toy{Name: "Ball"}})
	//db.Create(&Dog{Name: "Jerry", Toy: Toy{Name: "String"}})
	//db.Create(&Cat{Name: "Jenny", Toy: Toy{Name: "String"}})

	var dogs []Dog
	var cats []Cat
	db.Debug().Preload("Toy").Find(&dogs)
	fmt.Println(dogs)
	db.Debug().Preload("Toy").First(&cats)
}
