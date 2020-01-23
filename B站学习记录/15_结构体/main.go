package main

// 和TS一样一样的

import "fmt"

type person2 struct {
	name string
	age  int
}
type dog struct {
	name string
}

func main() {
	// typeDemo()
	// structDemo1()
	// structDemo2()
	structDemo3()
}
func newPersion(name string, age int) *person2 {
	return &person2{
		name: name,
		age:  age,
	}
}

func newDog(name string) dog {
	return dog{
		name: name,
	}
}

func structDemo3() {
	//构造函数:约定成俗用new开头
	//返回的是结构体还是结构体指针
	//当结构体比较大的时候尽量使用结构体指针减少内存开销
	p1 := newPersion("Mark", 38)
	p2 := newPersion("Kayla", 27)
	fmt.Println(p1, p2)

	d1 := newDog("周林")
	fmt.Println(d1)

}

func f1(x *person2) {
	(*x).age = 67
	// x.age = 67 //语法糖,自动根据指针找对应变量
}

func structDemo2() {
	var p person2
	p.name = "Mark"
	p.age = 24
	f1(&p)
	fmt.Println(p)

	//船舰指针类型的结构体
	var p2 = new(person2)
	fmt.Printf("person2 的指针%T\n", p2)
	fmt.Println(p2)

}

func structDemo1() {
	type person struct {
		name   string
		age    int
		hobby  []string // 如果不赋值则为空值
		gender string
	}

	var p person

	p.name = "Mark"
	p.age = 26
	p.gender = "男"
	p.hobby = []string{"篮球", "足球", "双色球"}

	fmt.Printf("p类型%T \n p值%v \n", p, p)

	//匿名结构体
	var s struct {
		name string
		age  int
	}
	fmt.Printf("s类型%T \n s值%v \n", s, s)

}

func typeDemo() {

	/*
		自定义类型和类型别名:
		基于内置的类型造一个自己的类型

	*/

	// type 后面跟的是一个类型 , 自定义类型例子:

	/*

	   区别,:
	   自定义类型是自己的类型
	   类型别名只在编写过程中有效

	*/
	type myInt int

	//类型别名
	type yourInt = int
	var n myInt
	n = 100
	fmt.Printf("%T\n", n) // main.myInt

	var m yourInt
	m = 199
	fmt.Printf("%T\n", m)

	var c rune
	c = '中'
	fmt.Printf("%T\n", c)

}
