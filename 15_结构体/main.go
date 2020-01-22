package main

import "fmt"

func main() {
	// typeDemo()
}

func typeDemo() {

	/*
		自定义类型和类型别名:
		基于内置的类型造一个自己的类型

	*/

	// type 后面跟的是一个类型 , 自定义类型例子:

	/*

	   区别,:
	   自定义类型是自己的类型
	   类型别名只在编写过程中有效

	*/
	type myInt int

	//类型别名
	type yourInt = int
	var n myInt
	n = 100
	fmt.Printf("%T\n", n) // main.myInt

	var m yourInt
	m = 199
	fmt.Printf("%T\n", m)

	var c rune
	c = '中'
	fmt.Printf("%T\n", c)

}
