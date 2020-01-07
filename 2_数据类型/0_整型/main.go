package main

/*

	golang 无法直接定义二进制 , 只能先定义为十进制然后转化为二进制

*/

import "fmt"

func main() {
	//十进制
	var a int = 10
	fmt.Printf("%d\n", a) // 十进制
	fmt.Printf("%b\n", a) //二进制   占位符 %b 表示二进制
	fmt.Printf("%o\n", a) //八进制   占位符 %o 表示二进制
	fmt.Printf("%x\n", a) //十六进制   占位符 %x 表示二进制

	fmt.Println("----------")

	//八进制 以0开头
	var b int = 077
	fmt.Printf("%o\n", b) // 77go

	//十六进制  0x 开头为16进制
	var c int = 0xff
	fmt.Printf("%x\n", c) // ff

	fmt.Println("------变量类型检测----")

	fmt.Printf("%T\n", c) //查看变量的类型

	//声明int8类型变量
	i4 := int8(9)

	fmt.Printf("%T\n", i4) //查看变量的类型
}
