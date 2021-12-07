package main

import "fmt"

/*
1. init 函数：初始化函数，可以用来进行一些初始化的操作
2. 全局变量定义init 比 main函数更优先执行
3. 多个源文件都有init函数的时候，如何执行
		 按照导入顺序执行，main 最后执行
*/

var num int = test() // 第一步：全局变量的定义

func test() int {
	fmt.Println("test函数被执行")
	return 10
}

func init() { // 第二步：init 函数的调用

	fmt.Println("init 函数")

}

func main() { // 第三步： main 函数的调用

	fmt.Println("main 函数")
}
