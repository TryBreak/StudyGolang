package main

import (
	"fmt"

	"github.com/shopspring/decimal"
)

// go get github.com/shopspring/decimal

func main() {
	n1 := "0.00035"
	n2 := "0.00036"

	// 加法
	a := Add(n1, n2)
	fmt.Println("-----加法-----", a)
	b := Sub(n1, n2)
	fmt.Println("-----减法-----", b)
	c := Mul(n1, n2)
	fmt.Println("-----乘法-----", c)
	d := Div(n1, n2)
	fmt.Println("-----除法-----", d)
	e := Per(n1, n2) // n1/n2 * 100
	fmt.Println("-----百分比计算-----", e)

	// 计算价格上涨的百分比
	cl := Sub(n2, n1)
	c2 := Per(cl, n1)
	fmt.Println("----涨幅计算-----", c2)
}

func toDec(s string) decimal.Decimal {
	n, _ := decimal.NewFromString(s)
	return n
}

// 加法
func Add(n1, n2 string) string {
	n := toDec(n1).Add(toDec(n2))
	return n.String()
}

// 减法
func Sub(n1, n2 string) string {
	n := toDec(n1).Sub(toDec(n2))
	return n.String()
}

// 乘法
func Mul(n1, n2 string) string {
	n := toDec(n1).Mul(toDec(n2))
	return n.String()
}

// 除法
func Div(n1, n2 string) string {
	n := toDec(n1).Div(toDec(n2))
	return n.String()
}

// 百分比计算
func Per(n1, n2 string) string {
	decimal.DivisionPrecision = 4
	n := toDec(n1).Div(toDec(n2)).Mul(toDec("100"))
	return n.String()
}
