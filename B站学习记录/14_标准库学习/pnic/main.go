package main

import "fmt"

func f1() {
	fmt.Println(1)
}

func f2() {
	defer func() {
		recover() // 继续往后面走了
		// fmt.Println(err)
		fmt.Println("拯救一下")
	}()
	panic("出现了错误")
	fmt.Println(2)
}

func f3() {
	fmt.Println(3)
}

func main() {
	f1()
	f2()
	f3()
}
