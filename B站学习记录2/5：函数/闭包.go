package main

import "fmt"

// 函数功能：求和
// 函数的名字： getSum 参数为空
// getSum函数返回值为一个函数，这个函数的参数是一个int类型的参数，返回值也是int类型
func getSum() func(int) int {
	var sum int = 0
	fmt.Println("sum", sum)
	return func(num int) int {
		sum = sum + num
		return sum
	}
}

/*

闭包的本质：
闭包的本质依旧是一个匿名函数，只是这个函数引入外界的变量/参数，匿名函数 + 引用的变量/参数 = 闭包

闭包可以保留你上次引用的某一个值

通过闭包可以反复引用某一个值

*/

func main() {
	f := getSum()
	h := getSum()
	fmt.Println("闭包", f(1)) //2
	fmt.Println("闭包", f(1)) //2
	fmt.Println("闭包", h(2)) //2
	fmt.Println("闭包", h(2)) //4

}
