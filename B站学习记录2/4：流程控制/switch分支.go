package main

import "fmt"

func main() {
	var score uint = 87

	switch true {
	case score >= 90:
		fmt.Println("您的成绩为A级别")
	case score >= 80:
		fmt.Println("您的成绩为B级别")
		fallthrough
	case score >= 70:
		fmt.Println("您的成绩为C级别")
	case score >= 60:
		fmt.Println("您的成绩为D级别")
	case score < 60:
		fmt.Println("您的成绩为E级别")
	default:
		fmt.Println("成绩出异常")
	}

	/*
		switc 后是一个表达式 常量，变量，有返回值的函数
		case 后的数据类型必须跟 switch 保持一致
		case 多个表达式用逗号分隔 case 表达式1,表达式2
		fallthrough 关键字， case 语句后增加  fallthrough 关键字，会继续执行下一个case 叫做 switch 穿透

	*/
}
