package main

import "fmt"

func main() {
	// 定义二维数组

	var arr [2][3]int = [2][3]int{{1, 4, 7}, {2, 5, 8}}

	fmt.Println(arr)

	fmt.Printf("&arr 的地址为:%p \n", &arr)
	fmt.Printf("&arr[0] 的地址为:%p \n", &arr[0])
	fmt.Printf("&arr[0][0] 的地址为:%p \n", &arr[0][0])
	fmt.Printf("&arr[1] 的地址为:%p \n", &arr[1])
	fmt.Printf("&arr[1][0] 的地址为:%p \n", &arr[1][0])

}
