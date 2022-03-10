package main

import "fmt"

func main() {
	// arr := make([]int, 5, 10)

	// for i := 0; i < 400; i++ {
	// 	arr = append(arr, i)
	// 	fmt.Println("-------")
	// 	fmt.Println(arr)
	// }

	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}

	n := 6

	fmt.Println(arr[0:n])
	fmt.Println(arr[n:])

}
