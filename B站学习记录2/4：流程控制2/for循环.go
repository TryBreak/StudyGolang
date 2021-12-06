package main

import "fmt"

func main() {

	fmt.Println("for 循环")

	var sum int = 0
	for i := 1; i <= 5; i++ {
		fmt.Println("for 循环i", i)
		sum += i
		fmt.Println("for 循环sum", sum)
	}

}
