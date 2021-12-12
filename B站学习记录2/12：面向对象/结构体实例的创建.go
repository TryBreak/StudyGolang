package main

import "fmt"

type Teacher struct {
	Name   string
	Age    int
	School string
}

func main() {
	fmt.Println("---方式一-----")
	var t1 Teacher
	t1.Name = "马士兵"
	t1.Age = 45
	t1.School = "清华大学"

	fmt.Printf("ma 的类型是：%T,值：%v \n", t1, t1) //main.Teacher,值：{ 0 }

	fmt.Println("---方式二-----")

	var t2 Teacher = Teacher{"马士兵", 45, "清华大学"}
	// t1.Name = "马士兵"
	// t1.Age = 45
	// t1.School = "清华大学"
	fmt.Println("t2", t2, t2.School)

	fmt.Println("---方式三-----")
	// 返回结构体的指针
	var t3 *Teacher = new(Teacher)
	// t3 指向一个地址
	// 应该给这个地址的指向 的对象 的字段 赋值：

	(*t3).Name = "王五"
	(*t3).Age = 65
	t3.School = "机械学院"
	// go底层对 t3.School 转化 成了 （*t3）为了符合程序员编程习惯

	fmt.Println("t3", t3, t3.School)
	fmt.Printf("t3 的地址：%p \n", &t3)

	fmt.Println("---方式四-----")

	var t4 *Teacher = &Teacher{}
	// var t4 *Teacher = &Teacher{"王五2", 654, "机械学院3"}
	(*t4).Name = "王五2"
	(*t4).Age = 654
	t4.School = "机械学院3"
	fmt.Println("t4", t4, t4.School)

}
