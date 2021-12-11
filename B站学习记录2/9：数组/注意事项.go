package main

import "fmt"

func main() {
	fmt.Println("--数组长度是类型的一部分,不可改改变----")
	var arr1 = [3]int{9, 6, 5}
	fmt.Printf("数组的类型为:%T -- %v \n", arr1, arr1)

	var arr2 = [6]int{6, 6, 8, 4, 2, 5}
	fmt.Printf("数组的类型为:%T -- %v \n", arr2, arr2)

	var arr3 []int = []int{3, 9, 5}
	test(&arr3)
	fmt.Printf("arr3 数组的类型为:%T -- %v \n", arr3, arr3)

	var arr4 [9]int

	arr4[0] = 98
	arr4[1] = 98
	arr4[2] = 98
	arr4[3] = 98
	arr4[4] = 98

	fmt.Printf("arr4 -- %v \n", arr4)

}

func test(arr *[]int) {
	(*arr)[1] = 999
}
