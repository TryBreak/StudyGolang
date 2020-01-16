package main

import "fmt"

func f1() {
	fmt.Println("hello")
}

func f2() int {
	return 10
}

//函数当作参数
func f3(x func()) {
	x()
}

func main() {
	a := f1
	fmt.Printf("%T \n", a)
	b := f2
	fmt.Printf("%T \n", b)

	f3(f1)
}
