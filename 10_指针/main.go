package main

import "fmt"

/*
	区别于C/C++中的指针，Go语言中的指针不能进行偏移和运算，是安全指针。
	要搞明白Go语言中的指针需要先知道3个概念：指针地址、指针类型和指针取值。
*/
func main() {
	getDemo()
}

func getDemo() {
	/*
		1. &: 取地址
		2. *: 根据地址取值
	*/
	n := 18
	p := &n // 取地址
	fmt.Println(p)
	fmt.Printf("%T \n", p)
	//取值
	m := *p
	fmt.Println(m)
	fmt.Printf("%T \n", m) // *int  指的是 int 类型的指针

}
