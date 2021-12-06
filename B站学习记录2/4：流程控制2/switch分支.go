package main

import "fmt"

func main() {

	/* 定义局部变量 */
	/* 	var grade string = "B"
	   	var marks int = 90

	   	switch marks {
	   	case 90:
	   		grade = "A"
	   	case 80:
	   		grade = "B"
	   	case 50, 60, 70:
	   		grade = "C"
	   	default:
	   		grade = "D"
	   	}

	   	switch {
	   	case grade == "A":
	   		fmt.Printf("优秀!\n")
	   	case grade == "B", grade == "C":
	   		fmt.Printf("良好\n")
	   	case grade == "D":
	   		fmt.Printf("及格\n")
	   	case grade == "F":
	   		fmt.Printf("不及格\n")
	   	default:
	   		fmt.Printf("差\n")
	   	}
	   	fmt.Printf("你的等级是 %s\n", grade) */
	/*
		// switch 语句还可以被用于 type-switch 来判断某个 interface 变量中实际存储的变量类型。
			switch x.(type){
				case type:
					statement(s);
				case type:
					statement(s);
				// 你可以定义任意个数的case
				default: //可选
					statement(s);
			}

	*/

	/* 	var x interface{}

	   	switch i := x.(type) {
	   	case nil:
	   		fmt.Printf(" x 的类型 :%T \n", i)
	   	case int:
	   		fmt.Printf("x 是 int 型 \n")
	   	case float64:
	   		fmt.Printf("x 是 float64 型 \n")
	   	case func(int) float64:
	   		fmt.Printf("x 是 func(int) 型 \n")
	   	case bool, string:
	   		fmt.Printf("x 是 bool 或 string 型 \n")
	   	default:
	   		fmt.Printf("未知型")
	   	}
	*/
	// 1. switch 后是一个表达式（即常量值、变量、一个有返回值的函数等都可以）
	// 2. case 后的各个值的类型必须和 switch 的表达式数据类型一致
	var score int = 1
	switch score / 10 {
	case 10:
		fmt.Println("等级为A")
	case 9:
		fmt.Println("等级为A")
	case 8:
		fmt.Println("等级为B")
	case 7:
		fmt.Println("等级为C")
	case 6:
		fmt.Println("等级为D")
	case 5:
		fmt.Println("等级为E")
	case 4:
		fmt.Println("等级为E")
	case 3:
		fmt.Println("等级为E")
	case 2:
		fmt.Println("等级为E")
	case 1:
		fmt.Println("等级为E")
	case 0:
		fmt.Println("等级为A")
	default:
		fmt.Println("成绩有误")
	}

	// 3. case 后面可以带多个表达式，使用逗号间隔 case 表达式1，表达式2，表达式3

	// 4. case 后面的表达式如果是常量值（字面量），则要求不能重复
	var a int = 12
	switch {
	case a == 1:
		fmt.Println("您的等级为A")
	case a >= 12:
		fmt.Println("您的等级为c")
	}

	// 5. case 后面不需要带 break，程序匹配到一个case后就会执行对应的代码块，然后退出 switch，如果一个都匹配不到，则执行default

	switch b := 7; {
	case b > 6:
		fmt.Println("大于6")
	case b > 6:
		fmt.Println("不大于6")
	}

	// 6. default 语句不是必须的

	// 7. switch 后也可以不带表达式，当做if分支来使用

	// 8. switch 后也可以直接声明定义一个变量，分号结束，不推荐

	// 9. switch 穿透，利用 fallthrough 关键字，如果在 case 语句块后增加 fallthrough，则会继续执行下一个case，这也叫做 switch 穿透

}
