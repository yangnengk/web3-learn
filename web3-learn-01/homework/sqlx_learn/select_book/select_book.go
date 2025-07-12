package selectbook

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Book struct {
	ID     int
	Price  float64
	Title  string
	Author string
}

func SelectBook(db *sqlx.DB) {

	// db.MustExec("insert into books(title,author,price) values(?,?,?)", "西游记", "吴承恩", 100.00)
	// db.MustExec("insert into books(title,author,price) values(?,?,?)", "水浒传", "施耐庵", 120.00)
	// db.MustExec("insert into books(title,author,price) values(?,?,?)", "三国演义", "罗贯中", 150.00)
	// db.MustExec("insert into books(title,author,price) values(?,?,?)", "红楼梦", "曹雪芹", 200.00)

	querySql := "select * from books where price > ?"

	var books []Book
	error := db.Select(&books, querySql, 50)
	if error != nil {
		panic(error)
	}
	fmt.Println(books)
}
