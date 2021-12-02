package main

import (
	"fmt"
)

func main() {
	fmt.Println("---通过指针改变值---")

	var num int64 = 10
	fmt.Println(num)

	num = 987
	fmt.Println(num)

	fmt.Println(&num)

	var ptr *int64 = &num // 地质类型的值一定是地址。地址的类型一定是匹配的
	*ptr = 99
	fmt.Println(num)

	fmt.Println("---指针变量接受的一定是地址值---")

}
