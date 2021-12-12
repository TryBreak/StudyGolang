package main

import "fmt"

type Student struct {
	Age int
}

type Person struct {
	Age int
}

type Stu Student

func main() {
	var s Student = Student{10}
	var p Person = Person{20}
	s = Student(p)
	// 一个结构体相当于一个类型

	fmt.Printf("s 的类型是：%T --- %v \n", s, s)
	fmt.Printf("p 的类型是：%T --- %v \n", p, p)

}
