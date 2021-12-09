package main

import (
	"fmt"
)

func main() {

	fmt.Println("内置函数", len("asdasd"))
	/*
		内置函数就是不需要导入包的函数
	*/
	fmt.Println("---new 分配内存-----")
	// 返回该内存新分配的零值的指针

	var n1 *int = new(int) // 参数是一个类型
	fmt.Println("n1", n1)

}
