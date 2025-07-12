package selectlearn

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Employee struct {
	ID         int     `db:"id"`
	Name       string  `db:"name"`
	Department string  `db:"department"`
	Salary     float64 `db:"salary"`
}

func Run(db *sqlx.DB) {
	// db.MustExec("insert into employees(name,department,salary) values(?,?,?)", "王五", "技术部", 10000)
	// db.MustExec("insert into employees(name,department,salary) values(?,?,?)", "赵六", "市场部", 12000)

	queryMultiple := "select * from employees where department = ?"
	var employees []Employee
	error := db.Select(&employees, queryMultiple, "技术部")
	if error != nil {
		panic(error)
	}
	fmt.Println(employees)
}

func RunQueryMax(db *sqlx.DB) {
	queryMax := `SELECT 
					e.id,
					e.salary,
					e.department,
					e.name
				FROM (
					SELECT 
						e.id,
						e.salary,
						e.department,
						e.name,
						ROW_NUMBER() OVER(ORDER BY e.salary DESC) r 
					from employees e
				) e WHERE e.r=1`
	var employees []Employee
	error := db.Select(&employees, queryMax)
	if error != nil {
		panic(error)
	}
	fmt.Println(employees)
}
