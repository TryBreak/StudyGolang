package main

import "fmt"

func main() {
	fmt.Println("Hello World1")
	fmt.Println("Hello World2")
	fmt.Println("Hello World3")
	var age = 19
	fmt.Println(age)
	fmt.Println("asdasdasd",
		"asdasdasd", "asdasdasd", "asdasdasd", "asdasdasd", "asdasdasd", "asdasdasd")

	var s3 string = "爱我a中华"
	var r3 = []rune(s3)
	fmt.Println((r3))
	fmt.Println(string(r3[0]))
	fmt.Println(s3[1])
	fmt.Println(s3[2])
	fmt.Println(s3[3])
	fmt.Println(s3[4])
}
