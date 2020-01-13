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
}

// GO 中没有more参数这个概念 , 所见即所得

func main() {
	// fmt.Println(sum(10, 20))
	// fmt.Println(f5(4, 5))

	// _, n := f6()
	// fmt.Println(n)
	f7("阿松大", 2, 3, 4)
}
