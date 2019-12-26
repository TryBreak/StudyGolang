package main

import "fmt"

//变量声明
var name1 string
var age1 int

// 批量声明
var (
	name string // ""
	age  int    // 0
	isOk bool   // false
)

//推荐使用小驼峰命名法

// var StudentName string //会爆黄色警告

// var student_name string //会爆黄色警告

var studentName string // 全局可声明变量且不使用

func main() {
	name = "理想"
	age = 18
	isOk = true
	fmt.Printf("name11:%s", name) // %s 是占位符 , 用 `name` 去替换

	var name string = "nnnnme" // 函数内的变量声明优先级最高

	var heiheihei string //函数作用域内声明而不使用会报错

	fmt.Print(isOk)               // 终端输出打印内容 不换行
	fmt.Println()                 // 打印换行
	fmt.Println("age:", age)      // 打印完之后加一个换行符
	fmt.Printf("name22:%s", name) // %s 是占位符 , 用 `name` 去替换

	heiheihei = "Mark"

	fmt.Println(heiheihei)

	// 声明变量同时赋值

	var si = "mark" //无需声明类型 , 类型推导
	fmt.Println(si)

	//简短变量声明 , 只能在函数作用域内使用
	s1 := "123"
	fmt.Println(s1)

	// 匿名变量 用单独一个`_` 来表示 , 函数接收但不使用的时候
	// 匿名变量不占用命名空间，不会分配内存，
	//
	/**
	匿名变量不占用命名空间，不会分配内存，
	所以匿名变量之间不存在重复声明。 (在Lua等编程语言里，匿名变量也被叫做哑元变量。)
		注意事项：
			函数外的每个语句都必须以关键字开始（var、const、func等）
			`:=` 不能使用在函数外。
			`_` 多用于占位，表示忽略值。
	*/

}
