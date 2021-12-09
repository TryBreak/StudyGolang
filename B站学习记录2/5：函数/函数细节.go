package main

import "fmt"

func 加法(数字1 int, 数字2 int) int {
	var 返回数 int = 0
	返回数 += 数字1
	返回数 += 数字2
	return 返回数
}

// 定义一个函数,可变参数(参数的数量可变)
// args...int 可以传入任意多个数量的 int 类型的数据 掺入 0个 1个 ...n个
func test(args ...int) {
	// 函数内部处理可变参数，当作切片来处理
	for i, v := range args {
		fmt.Println(i, v)
	}

}

func main() {
	var 值1 int = 加法(10, 20)
	fmt.Println(值1)

	// golang 的函数不支持重载
	// 函数名相同，但是参数不同
	// 函数支持可变参数
	// 函数内修改不会影响原来的值
	fmt.Println("--start-----")
	test()
	fmt.Println("-------")
	test(3)
	fmt.Println("-------")
	test(37, 38, 39)
	fmt.Println("-------")

}
