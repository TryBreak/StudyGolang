package main

import (
	"fmt"
)

func main() {

	// 1 变量的声明
	// var age int

	// // 2 变量的赋值
	// age = 278

	// // 3 变量的使用
	// fmt.Println(age)

	// 变量的 4 种使用形式

	// var age float64 = -0.123141231

	// fmt.Println(age)

	/* var num1 float64 = 3.14
	fmt.Println(num1)

	var num2 float64 = -3.14
	fmt.Println(num2)

	var num3 float32 = 314e-2
	fmt.Println(num3)

	var num4 float64 = 314e+2
	fmt.Println(num4)

	var num5 float32 = 314e+2
	fmt.Println(num5)

	var num6 float32 = 256.0000000000916
	fmt.Println(num6)

	var num7 = 256.0000000000916
	fmt.Println(num7)
	fmt.Println("num7的类型是： ") */

	var s1 uint = '中'
	fmt.Println(s1)                  // unicode 码值
	fmt.Printf("s1对应的字符为：%c \n", s1) // unicode 码值

	var s2 = "阿斯达中"
	var s3 = '中'
	fmt.Println(s2) // 字符串
	fmt.Println(s3) // unicode 码值

	/*

		字母，数字，标点, ASCII 码 是  unicode 符集的前 128位

		汉字字符是 unicode 码

		golang 的字符实际上对应的是 UTF-8 是 unicode 其中一种编码方案

	*/

}
