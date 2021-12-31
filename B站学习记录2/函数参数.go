package main

import "log"

func fun2(s string) {
	log.Println(s)
}

func main() {
	func1(fun2)
}

func func1(f func(string)) {
	s1 := "asd"
	f(s1)
}
