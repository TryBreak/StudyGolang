package main

import (
	"fmt"
	"math"
)

/*
Go语言中只有强制类型转换，没有隐式类型转换。该语法只能在两个类型之间支持相互转换的时候使用。


基本语法:


T(表达式)

其中，T表示要转换的类型。表达式包括变量、复杂算子和函数返回值等.

*/

func main() {
	var (
		a = 3
		b = 4
	)
	var c int

	x := math.Sqrt(float64(a*a + b*b))

	c = int(x)

	fmt.Printf("%T \n", x)
	fmt.Printf("%T \n", c)
}
