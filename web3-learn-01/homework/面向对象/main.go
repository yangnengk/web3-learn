package main

import (
	"fmt"
)

func main() {
	var r = Rectangle{
		Width:  10,
		Height: 20,
	}
	var c = Circle{
		Radius: 10,
	}

	fmt.Println("方形面积：", r.Area())
	fmt.Println("方形周长：", r.Perimeter())
	fmt.Println("圆的面积：", c.Area())
	fmt.Println("圆的周长：", c.Perimeter())

	var e = Employee{
		Person: Person{
			Name: "张三",
			Age:  20,
		},
		EmployeeID: "123456",
	}
	e.Person.Name = "李四"
	e.PrintInfo()
}

type Shape interface {
	Area()
	Perimeter()
}

type Rectangle struct {
	Shape
	Width  float64
	Height float64
}

type Circle struct {
	Shape
	Radius float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.Radius
}

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	EmployeeID string
}

func (e Employee) PrintInfo() {
	fmt.Printf("员工的信息，姓名：%s， 员工的年龄%d， 员工id：%s", e.Name, e.Age, e.EmployeeID)
}
