package main

import "fmt"

func test(num int) {
	fmt.Println(num)
}

func main() {

	a := test
	fmt.Printf("a的类型是：%T，test函数的类型是：%T", a, test)
}
