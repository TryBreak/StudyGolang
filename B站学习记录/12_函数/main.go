package main

import "fmt"

//函数的定义
func sum(x int, y int) (xxx int) {
	return x + y
}

//没有返回值的
func f1(x int, y int) {
	fmt.Println(x + y)
}

//没有参数没有返回值
func f3() {
	fmt.Println("f2")
}

// 返回值可以命名也可以不命名
// 没有参数但有返回值
func f4() int {
	ret := 3
	return ret
}

func f5(x int, y int) (ret int) {
	ret = x + y
	return // 使用命名返回值 return 可以省略
}

//多个返回值
//当参数连续两个参数的类型一致时 , 我们可以将非最后一个参数的类型省略
func f6(x, y int, m, n string, i, j bool) int {
	return x + y
}

// 可变长度的参数

func f7(x string, y ...int) {
	fmt.Println(x)
	fmt.Println(y) //y的类型是切片

	//不可以在命名函数中声明命名函数
	// func f8()  {
	// 	return 9
	// }

	//但是可以 声明匿名函数

	func() {
		y[0]++
	}()

}

// GO 中没有more参数这个概念 , 所见即所得

func main() {
	// fmt.Println(sum(10, 20))
	// fmt.Println(f5(4, 5))

	// _, n := f6()
	// fmt.Println(n)
	// f7("阿松大", 2, 3, 4)

	//defer

	// deferDemo()

	// f111()
	// defer f222()
	// defer f333() // 多个defer 按照先进后出(后进先出)的顺序执行
	// f444()

	// s := []int{1, 2, 3, 4, 5}
	// fmt.Println(s)

	// fff()

	x := f98()
	fmt.Println(x)

}

func deferDemo() {
	fmt.Println("start")
	defer fmt.Println("嘿嘿嘿") // 延迟到函数即将返回的时候执行
	fmt.Println("end")
}

func f111() {
	fmt.Println(1111111111111)
}
func f222() {
	fmt.Println(22222222222)
}
func f333() {
	fmt.Println(333333333333)
}
func f444() {
	fmt.Println(444444444444)
}

var fff = func() {
	fmt.Println("sadsadsad")
}

/*
go 中的 return不是原子操作

分为两步执行

1. 返回值赋值
2. 真正的RET返回

函数中如果存在defer , 那么 defer 则自第一步和第二步之间

*/

func f98() (aa int) {
	x := 5
	aa = x
	defer func() {
		x++
	}()
	return
}
