package main

import "fmt"

func test(num int) {
	fmt.Println(num)
}

// 函数也是一个类型
type myFunc func(int) // 使用 type 关键词来自定义类型

func test02(num1 int, num2 int, testFunc func(int)) {
	fmt.Println(num1, num2)
}

// 交换两属性的值

func exchangeValue(num1 *int, num2 *int) {
	var t int
	t = *num1
	*num1 = *num2
	*num2 = t
}

type myInt int

func test05(num1 int, num2 int) (sum int, sub int) {
	sum = num1 - num2
	sub = num2 - num1
	return
}

func main() {

	a := test
	fmt.Printf("a的类型是：%T，test函数的类型是：%T ", a, test)
	a(60)

	var t1 = 30
	var t2 = 10
	fmt.Printf("t1 的值是：%v，t2的值是：%v \n", t1, t2)
	exchangeValue(&t1, &t2)
	fmt.Printf("t1 的值是：%v，t2的值是：%v \n", t1, t2)

	fmt.Println("------函数别名-------")
	test02(10, 20, a)

	fmt.Println("------自定义数据类型-------")

	var num1 myInt = 30
	fmt.Printf("num1 的值是：%v，类型是：%T \n", num1, num1)

	fmt.Println("------对返回值进行命名-------")

	var b, c = test05(10, 20)
	fmt.Printf("b 的值是：%v,c 的值是 %v \n", b, c)

	func test06(num1 int, num2 int) (sum int, sub int) {
		sum = num1 - num2
		sub = num2 - num1
		return
	}

}
