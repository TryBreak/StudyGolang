package main

import "fmt"

/*
	区别于C/C++中的指针，Go语言中的指针不能进行偏移和运算，是安全指针。
	要搞明白Go语言中的指针需要先知道3个概念：指针地址、指针类型和指针取值。
*/

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

func newDemo() {
	//new关键字折你去哪个内存地址
	var p *int // nil pointer
	fmt.Printf("p --- %v \n", p)

	var a = new(int) //new 关键字申请内存地址
	fmt.Printf("a --- %v \n", a)
	fmt.Printf("*a --- %v \n", *a) // 取值  - 0

	*a = 100

	fmt.Printf("*a --- %v \n", *a) // 向该内存地址赋值100

	var s = new(string)
	fmt.Printf("s --- %v \n", s)
	fmt.Printf("*s --- %v \n", *s) // 取值  - 空
	*s = "beizou"

	fmt.Printf("*s --- %v \n", *s) // 向该内存地址赋值100
}

func makeDemo() {
	//make也用于分配内存 它只用于
}
func main() {
	// getDemo()
	newDemo()
	// makeDemo()
}
