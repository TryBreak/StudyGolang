/*
切片 slice
是golang中特有的数据类型

数组有特定的用处，但是却有一些呆板（长度固定不变），所以在go语言中
并不是特别常见。

相对的，切片确实随处可见的，切片是一种建立在数组类型之上的抽象，

它构建在数组之上并且提供更强大的能力和便捷

切片是对数组一个连续片段的引用，所以切片是一个引用类型。
这个片段可以说是整个数组，或者是由起始和终止索引标识的
一些项的子集。需要注意的是，终止索引标识的项不包括在切片内。
切片提供了一个相关数组的动态窗口。

*/

package main

import "fmt"

func main() {
	fmt.Println("--------切片--------")
	var intarr [6]int = [6]int{3, 6, 9, 8, 5, 5}
	// 切片定义在数组之上
	var slice1 []int = intarr[1:3]
	// 下标1 到 下标2 ，含前不含后

	fmt.Printf("slice1 的长度是：%v , 切片的类型是：%T \n",
		len(slice1), slice1)

	fmt.Println("容量为", cap(slice1)) // 容量可以动态变化
	//len 为 2 ，cap 为两倍
	fmt.Println("slice1", slice1)
}
