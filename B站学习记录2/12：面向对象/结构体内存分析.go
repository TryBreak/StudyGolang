package main

import "fmt"

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

	/*
		内存分析：
		t1 值类型 ，开辟内存空间，里面包含属性
		没有赋值则表示没有值，不占用内存空间
		键值对的方式存在
	*/

}
