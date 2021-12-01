package main

import "fmt"

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
		n1 int     = 199
		n2 float64 = 3.14159265358979323846
		n3 bool    = false
		n4 byte    = 'a'
	)

	// 转成字符串
	var s1 = fmt.Sprintf("%d", n1)

	fmt.Printf("s1的类型是：%T , s1 = %q \n", s1, s1)

	var s2 = fmt.Sprintf("%f", n2)
	fmt.Printf("s2的类型是：%T , s2 = %q \n", s2, s2)

	var s3 = fmt.Sprintf("%t", n3)
	fmt.Printf("s3的类型是：%T , s3 = %q \n", s3, s3)

	var s4 = fmt.Sprintf("%d", n4)
	fmt.Printf("s4的类型是：%T , s4 = %q \n", s4, s4)

}
