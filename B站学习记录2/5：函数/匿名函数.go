package main

import (
	"fmt"
	"testing"
)

var test01 = func(t *testing.T) {

}

func main() {

	// main 里面定义的函数，没有名字的函数
	result := func(num1 int, num2 int) int {
		return num1 + num2
	}
	fmt.Println("匿名函数", result(10, 20))
	test01()

}
