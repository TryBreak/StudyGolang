package alse

import "fmt"

//Add is a func
func Add(a, b int) int {
	fmt.Println("调用Add函数")
	return a * b
}
