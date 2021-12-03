package main

import (
	"fmt"
	test "test/test"
)

func main() {
	var age int64 = 64
	fmt.Println("这里是main", age)
	fmt.Println("这里是test下的变量", test.StuNo)
}
