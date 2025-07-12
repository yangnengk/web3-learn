package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	selectbook "github.com/learn/sqlx_learn/select_book"
	selectlearn "github.com/learn/sqlx_learn/select_learn"
)

func main() {

	dsn := "root:123456@tcp(127.0.0.1:3306)/sqlx?charset=utf8mb4&parseTime=True&loc=Local"

	db, error := sqlx.Connect("mysql", dsn)
	if error != nil {
		panic(error)
	}

	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)

	selectlearn.Run(db)
	// selectlearn.RunQueryMax(db)

	selectbook.SelectBook(db)

	defer db.Close()
}
