package main

import "fmt"

/*
	和js的差不多
*/

var x = 100

func f1() {
	x := 10
	fmt.Println(x)
}

func main() {
	f1()
}
