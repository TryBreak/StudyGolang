package main

import (
	"fmt"
	"strconv"
)

func main() {
	/* var n1 int = 100
	// var n2 uint = n1
	fmt.Println(n1)

	var n2 float64 = float64(n1)
	fmt.Println(n2)
	fmt.Printf("n2的类型： %T \n", n2)
	fmt.Printf("n1的类型： %T \n", n1)

	var n3 int64 = 128
	var n4 int8 = int8(n3) // -128

	fmt.Printf("n3的类型： %T \n", n3)
	fmt.Printf("n4的类型： %T \n", n4)
	fmt.Printf("n3的值：%v \n", n3)
	fmt.Printf("n4的值：%v \n", n4) */

	var (
	// n1 int     = 199
	// n2 float64 = 3.14159265358979323846
	// n3 bool    = false
	// n4 byte    = 'a'
	)

	// 转成字符串
	// var s1 = fmt.Sprintf("%d", n1)

	// fmt.Printf("s1的类型是：%T , s1 = %q \n", s1, s1)

	// var s2 = fmt.Sprintf("%f", n2)
	// fmt.Printf("s2的类型是：%T , s2 = %q \n", s2, s2)

	// var s3 = fmt.Sprintf("%t", n3)
	// fmt.Printf("s3的类型是：%T , s3 = %q \n", s3, s3)

	// var s4 = fmt.Sprintf("%d", n4)
	// fmt.Printf("s4的类型是：%T , s4 = %q \n", s4, s4)

	// strconv
	// strconv包实现了基本数据类型和其字符串表示的相互转换。

	var n1 int = 199
	var s1 string = strconv.FormatInt(int64(n1), 10)
	// 第一个参数必须是 int64 ， 第二个参数是字面量十进制
	fmt.Printf("s1 的类型是：%T , s1 = %q \n", s1, s1)

	var n2 float64 = 3.14159265358979323846
	var s2 string = strconv.FormatFloat(float64(n2), 'f', 20, 64)
	var s3 string = fmt.Sprintf("%f", n2)
	var s4 string = strconv.FormatFloat(float64(n2), 'f', 20, 32)
	// 第一个参数是 float64 第二个参数是格式 ‘f’ 为十进制，第三个参数是精度控制(小数点后多少位)，第四个参数是来源类型
	fmt.Printf("n2 的类型是：%T , n2 = %v \n", n2, n2)
	fmt.Println(n2)
	fmt.Printf("s2 的类型是：%T , s2 = %q \n", s2, s2)
	fmt.Printf("s4 的类型是：%T , s4 = %q \n", s4, s4)
	fmt.Printf("s3 的类型是：%T , s3 = %q \n", s3, s3)
	// big 的库

	var n5 bool = true
	var s5 string = strconv.FormatBool(n5)
	fmt.Printf("s5 的类型是：%T , s5 = %v \n", s5, s5)
	fmt.Println(s5)

	// 将字符串转为基础类型

	// var s1 string = "true"
	// var b bool
	// b, _ = strconv.ParseBool(s1)
	// fmt.Printf("b 的类型是： %T，b = %v \n", b, b)

	// var s2 string = "256"
	// var f1 int64
	// f1, _ = strconv.ParseInt(s2, 10, 64)
	// fmt.Printf("f1 的类型是： %T，f1 = %v \n", f1, f1)

}
