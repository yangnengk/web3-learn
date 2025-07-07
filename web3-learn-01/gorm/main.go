package main

import (
	"github.com/learn/gorm/lesson03"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, error := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"))
	if error != nil {
		panic(error)
	}
	// lesson01.Run(db)
	// lesson02.Run(db)
	lesson03.Run(db)

}
