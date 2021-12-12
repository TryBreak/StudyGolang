package main

import "fmt"

// 结构体和结构体面向对象的引入

// fmt.Println("---定义一个老师的结构体---")

type Teacher struct {
	Name   string
	Age    int
	School string
}

func main() {
	// 创建老师结构体的实例
	var t1 Teacher

	fmt.Printf("ma 的类型是：%T,值：%v \n", t1, t1) //main.Teacher,值：{ 0 }

	t1.Name = "马士兵"
	t1.Age = 45
	t1.School = "清华大学"

	fmt.Printf("t1 的类型是：%T,值：%v \n", t1, t1)
	fmt.Println("t1 的名字是：", t1.Name)

}
