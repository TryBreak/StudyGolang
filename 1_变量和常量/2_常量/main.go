package main

import "fmt"

// 常量
const pi = 3.1415926

//批量声明常量
const (
	statusOkk = 200
	notFund   = 404
)

// 批量声明常量时 , 如果某一行声明后没有赋值 , 则默认跟上一行一样
const (
	n1 = 100
	n2
	n3
)

func main() {
	fmt.Println("n1:", n1)
	fmt.Println("n2:", n2)
	fmt.Println("n3:", n3)
}
