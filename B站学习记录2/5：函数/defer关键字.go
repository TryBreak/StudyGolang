package main

// defer 场景

/*

 */

import "fmt"

func main() {

	add(10, 20)

}

func add(num1 int, num2 int) int {
	/*
		在golang中，程序遇到defer关键字，不会立即执行defer后的语句，
		而是将defer后的语句压入一个栈当中，然后继续执行函数后面的内容
	*/
	defer fmt.Println("num1=", num1)
	defer fmt.Println("num2=", num2)
	// 先进后出
	num1 += 100
	num2 += 100

	var sum int = num1 + num2
	fmt.Println("sum=", sum)
	return sum
}
