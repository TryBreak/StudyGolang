package main

import "fmt"

func main() {

	fmt.Println("------切片的内存分析------------")

	var arrint [6]int = [6]int{3, 6, 9, 1, 4, 7}

	var slice1 []int = arrint[1:3]

	fmt.Println("arrint", arrint)
	fmt.Println("slice", slice1)

	fmt.Println("slice的元素个数", len(slice1))
	fmt.Println("slice的容量", cap(slice1))

	fmt.Printf("数组下标位置为1位置的地址：%p,值为：%v \n",
		&arrint[1], arrint[1])
	fmt.Printf("切片下标位置为1位置的地址：%p,值为：%v \n",
		&slice1[0], slice1[0])

	// 地址一致,  切片为引用类型,用的是同一块内存地址
}
