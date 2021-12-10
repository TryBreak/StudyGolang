package main

import "fmt"

func main() {
	var arr [3]int

	fmt.Printf("arr 的内存地址为：%p \n", &arr) // 和 &arr[0]  一样
	fmt.Printf("arr 的内存地址为：%p \n", &arr[0])
	fmt.Printf("arr 的内存地址为：%p \n", &arr[1])
	fmt.Printf("arr 的内存地址为：%p \n", &arr[2])

	arr[0] = 10
	arr[1] = 20
	arr[2] = 30

	fmt.Println("-----------------------")

	fmt.Printf("arr 的内存地址为：%p \n", &arr) // 和 &arr[0]  一样
	fmt.Printf("arr 的内存地址为：%p \n", &arr[0])
	fmt.Printf("arr 的内存地址为：%p \n", &arr[1])
	fmt.Printf("arr 的内存地址为：%p \n", &arr[2])
	// 地址可以算的到，地址的下标从零开始的
	// 数组的优点 是读取、查询、访问速度比较快
	// 因为其地址是连续的，每个空间占据的字节数，取决于数组类型
}
