package main

import "fmt"

func main() {
	var arr1 [5]int
	arr2 := [5]int{2, 4, 6}
	// arr3 := [...]int{1, 3, 5, 7, 9, 11}
	// var grid [4][5]int
	// fmt.Println(arr1, arr2, arr3)
	// fmt.Println(grid)

	/* 		for i, v := range arr3 {
		fmt.Println(i, v)
	} */
	// maxValueRule() // 求最大值
	printArray(arr1)
	printArray(arr2)
	// printArray(arr3) // 长度即是类型

	fmt.Println(arr1, arr2)

}

func printArray(arr [5]int) { //函数传值永远都是值类型
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func maxValueRule() {
	arr := [...]int{6, 4, 2, 7, 8, 2, 45, 7}

	maxi := -1
	maxValue := -1

	for i, v := range arr {
		if v > maxValue {
			maxi, maxValue = i, v
		}
	}
	fmt.Println(maxi, maxValue)
	// _ 可省略变量
}
