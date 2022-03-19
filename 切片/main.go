package main

import "fmt"

var L = []int{1, 2, 3, 4, 5, 6, 7, 8}

func main() {

	period := 3

	period_list := make([]int, period)

	copy(period_list, L)

	fmt.Println(period_list)
	fmt.Println(L)

}
