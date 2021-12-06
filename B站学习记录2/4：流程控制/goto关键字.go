package main

import "fmt"

func main() {

	fmt.Println("hello -- 1")
	fmt.Println("hello -- 2")
	if 1 == 1 {
		goto label1
	}

	fmt.Println("hello -- 3")
	fmt.Println("hello -- 4")
	fmt.Println("hello -- 5")
	fmt.Println("hello -- 6")

label1:
	fmt.Println("hello -- 7")
	fmt.Println("hello -- 8")

}
