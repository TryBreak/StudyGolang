package main

import "fmt"

func fibonacci() func() int {
	a := 0
	b := 1
	// a, b := 0, 1
	fmt.Println("a,b", a, b) //1

	return func() int {
		b = a + b
		a = b - a
		// a, b = b, a+b
		fmt.Println("2 - a,b", a, b) //1
		return a
	}
}

func main() {
	f := fibonacci()

	fmt.Println("f()", f()) //1
	fmt.Println("f()", f()) //1
	fmt.Println("f()", f()) //2
	fmt.Println("f()", f()) //3
	fmt.Println("f()", f()) //5
	fmt.Println("f()", f()) //8
	fmt.Println("f()", f()) //13
	fmt.Println("f()", f()) //21
}
