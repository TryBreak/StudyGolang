/*
函数的定义与调用
*/

package main

import (
	"fmt"
	"strconv"
)

/*

func 函数名(形参列表)(返回值类型){
	执行语句..
	return + 返回值列表
}

*/

func cal(num1 int, num2 int) int {
	var sum int = 0
	sum += num1
	sum += num2
	return sum
}

func cal2(num1 int, num2 int) (string, int) {
	var sum int = 0
	sum += num1
	sum += num2

	var result int = num1 - num2
	return strconv.FormatInt(int64(sum), 10), result
}

func main() {
	// 调用函数

	var sum = cal(10, 20)

	fmt.Println(sum)

	var sum2, result = cal2(10, 20)

	fmt.Printf("这个值的类型是 %T，值为%v \n", sum2, sum2)
	fmt.Printf("这个值的类型是 %T，值为%v \n", result, result)

}
