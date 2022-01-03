package main

import "fmt"

func main() {
	arr := make([]int, 5, 10)

	for i := 0; i < 400; i++ {
		arr = append(arr, i)
		fmt.Println("-------")
		fmt.Println(arr)
	}
}
