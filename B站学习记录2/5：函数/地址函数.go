package main

import "fmt"

func test(num *int) {
	*num = 30
}

func main() {

	var num int = 10
	test(&num)
	fmt.Println(num)

	// 函数本身也是一种数据类型
}
