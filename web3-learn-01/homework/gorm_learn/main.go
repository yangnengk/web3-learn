package main

import (
	transactiondemo "github.com/learn/gorm_learn/transaction_demo"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, error := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"))
	if error != nil {
		panic(error)
	}

	println("数据库连接成功")

	// 基本CRUD操作
	// cruddemo.Run(db)

	// 事务操作
	transactiondemo.Run(db)
}
