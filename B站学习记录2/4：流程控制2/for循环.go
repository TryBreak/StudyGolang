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
	fmt.Println("-----------")

	// for i
	var str string = "hello golang"

	fmt.Println(str[1])

	for i := 0; i <= len(str); i++ {
		fmt.Println(str[i])
	}

	fmt.Println("-----------")

	// for range
	var str2 string = "hello golang"

	for i, v := range str2 {
		fmt.Println(i, v)
	}

	fmt.Println("-----------")
}
