package main

import "fmt"

func main() {

	// 方式1:定义数组

	var arrint [6]int = [6]int{1, 2, 3, 4, 5, 6}
	var slice3 = arrint[1:3]
	fmt.Println("slice3", slice3)

	// 方式2:
	// make 去定义
	// 底层创建一个数组,对外部可见,所以不可以直接操作数组,
	// 要通过slice去间接访问各个元素,不可以通过数组进行维护
	var slice1 []int = make([]int, 4, 20)
	// 三个参数, 切片类型,切片长度,切片容量

	fmt.Println("slice1", slice1)

	fmt.Printf("slice1 的长度是:%v,类型是 %T,容量是:%v \n",
		len(slice1), slice1, cap(slice1))

	slice1[0] = 5
	slice1[3] = 8
	fmt.Println("slice1", slice1)

	// 方式3
	// 定一个切片,直接就指定具体数组,使用方式类似于 make

	var slice4 = [3]int{1, 3, 9}
	fmt.Println("slice3 ", slice4)
	fmt.Println("slice3 cap", cap(slice4))
	fmt.Println("slice3 len", cap(slice4))

}
