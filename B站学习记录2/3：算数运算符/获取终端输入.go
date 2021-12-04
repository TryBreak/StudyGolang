package main

import "fmt"

/*

需要用到的 AI

Scanin

Scanf

*/

func main() {

	fmt.Println("获取终端输入内容,年龄,姓名,成绩,是否 VIP")

	/* 	var age int64
	   	fmt.Println("请录入学生的年龄:")
	   	fmt.Scanln(&age) // 参数为内存空间的地址

	   	var name string
	   	fmt.Println("请录入学生的姓名:")
	   	fmt.Scanln(&name) // 参数为内存空间的地址

	   	var score float64
	   	fmt.Println("请录入学生的成绩:")
	   	fmt.Scanln(&score) // 参数为内存空间的地址

	   	var isVip bool
	   	fmt.Println("是否 Vip:")
	   	fmt.Scanln(&isVip) // 参数为内存空间的地址

	   	fmt.Printf("学生的年龄为:%v,姓名为:%v,学生的成绩为:%v,是否为 Vip %v \n", age, name, score, isVip)
	*/
	var (
		age   int64
		name  string
		score float64
		isVIP bool
	)

	fmt.Println("请输入学生的年龄,姓名,成绩,是否Vip,空格隔开")
	fmt.Scanf("%v %v %v %v", &age, &name, &score, &isVIP)
	fmt.Printf("学生的年龄为:%v,姓名为:%v,学生的成绩为:%v,是否为 Vip %v \n", age, name, score, isVIP)

}
