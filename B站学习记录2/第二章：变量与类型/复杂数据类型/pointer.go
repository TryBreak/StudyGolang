package main

import (
	"fmt"
)

func main() {
	// 没有指针运算，通过指针找变量，用过变量找指针
	fmt.Println("指针")

	var age int64 = 18 // 在内存中开辟出一个空间，装了个18
	fmt.Println(age)

	// 这个空间的地址是什么呢？
	fmt.Println(&age) // 显示 age 的内存 地址

	// 指针变量

	var ptr *int64 = &age

	fmt.Printf("ptr 的类型是：%T，ptr=%v \n", ptr, ptr)
	fmt.Printf("age 的类型是：%T age=%v \n", age, age)
	fmt.Printf("&age 的类型是：%T &age=%v \n", &age, &age)
	fmt.Printf("&ptr 的类型是：%T &ptr=%v \n", &ptr, &ptr)
	// 指针就是内存地址

	// 通过内存地址拿到内存当中的值
	fmt.Printf("*ptr 的类型是：%T *ptr=%v \n", ptr, *ptr)
	// 取地址当中的值  *
	// 获取内存的地址 &

}
