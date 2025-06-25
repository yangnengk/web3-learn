package main

import "fmt"

/*
interface接口
在 Go 中接口是一种抽象类型，是一组方法的集合，里面只声明方法，而没有任何数据成员。
而在 Go 中实现一个接口也不需要显式的声明，只需要其他类型实现了接口中所有的方法，就是实现了这个接口
*/
func main() {
	creditCard := &CreditCard{balance: 0, limit: 1000}
	debitCard := &DebitCard{balance: 500}

	fmt.Println("信用卡购买商品")
	purchaseItem(creditCard, 800)

	fmt.Println("借记卡购买商品")
	purchaseItem(debitCard, 300)

	fmt.Println("再次使用借记卡购买商品")
	purchaseItem(debitCard, 300)

	// 结构体转成接口
	var accountA Account = debitCard
	fmt.Println("借记卡账户A余额：", accountA.GetBalance())

	// 使用过程中需要注意以下几点：
	//1.Go 中接口声明的方法并不要求需要全部公开。
	//2.直接用接口类型作为变量时，赋值必须是类型的指针。
	//3.接口可以嵌套。
	//4.接口中声明的方法，参数可以没有名称。
	//5.如果函数参数使用 interface{}可以接受任何类型的实参。同样，可以接收任何类型的值也可以赋值给 interface{}类型的变量
	//var a interface{} = debitCard	// interface{}类型可以接受任何类型的值,相当于java
	//var a interface{} = 10
	//a = &creditCard // 可以赋值指针的指针
	//a = "aa"
	//a = *creditCard // 可以赋值为结构体
	fmt.Println("anyParam-----------------------------------")
	anyParam(creditCard)
	anyParam(debitCard)
	anyParam(10)
	anyParam(accountA)
	anyParam(&creditCard)
	anyParam(*creditCard)
}

// PayMethod  接口定义了支付方法的基本操作
type PayMethod interface {
	Account
	Pay(int) bool // 接口中声明的方法，参数可以没有名称
}

type Account interface {
	GetBalance() int
}

// CreditCard 结构体实现 PaymentMethod 接口
type CreditCard struct {
	balance int
	limit   int
}

func (c *CreditCard) Pay(amount int) bool {
	if c.balance+amount <= c.limit {
		c.balance += amount
		fmt.Printf("信用卡支付成功: %d \n", amount)
		return true
	}
	fmt.Printf("信用卡支付失败: 超出额度，余额：%d \n", c.balance)
	return false
}

func (c *CreditCard) GetBalance() int {
	return c.balance
}

// DebitCard 结构体实现 PaymentMethod 接口
type DebitCard struct {
	balance int
}

func (d *DebitCard) Pay(amount int) bool {
	if d.balance >= amount {
		d.balance -= amount
		fmt.Printf("借记卡支付成功: %d \n", amount)
		return true
	}
	fmt.Printf("借记卡支付失败: 余额不足，余额：%d \n", d.balance)
	return false
}

func (d *DebitCard) GetBalance() int {
	return d.balance
}

// 使用 PaymentMethod 接口的函数
func purchaseItem(p PayMethod, price int) {
	if p.Pay(price) {
		fmt.Printf("购买商品成功, 剩余余额： %d \n", p.GetBalance())
	} else {
		fmt.Printf("购买商品失败: %d \n", price)
	}
}

func anyParam(param interface{}) {
	fmt.Println("param is ", param)
}
