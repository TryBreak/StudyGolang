package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		fmt.Println(i)
		if i == 14 {
			return // 结束当前函数,当前循环
		}
	}

	fmt.Println("hello golang")
}
