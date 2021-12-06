package main

import "fmt"

func main() {
	fmt.Println("分支模式")
	// var count int = 100
	// if count < 30 {
	// 	fmt.Println("对不起，口罩存量不足")
	// } else {
	// 	fmt.Println("口罩数量充足")
	// }

	/*
		if 后面表达式，返回结果一定是 bool 类型
	*/

	/* 	if count := 80; count < 30 {
		fmt.Println("口罩数量不足")
	} */

	var score int = 72
	if score >= 90 {
		fmt.Println("您的成绩为A级别")
	} else if score >= 80 {
		fmt.Println("您的成绩为B级别")
	} else if score >= 70 {
		fmt.Println("您的成绩为C级别")
	} else if score >= 60 {
		fmt.Println("您的成绩为D级别")
	} else if score < 60 {
		fmt.Println("您的成绩为E级别")
	}
}
