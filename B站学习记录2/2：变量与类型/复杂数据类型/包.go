package main // 入口包

/*
	定义源文件， package main
	main 是一个程序的入口
	所以你main函数它所在的包建议定义为main包
	如果不定义为main包，那么就不能得到可执行文件


	package 包名要与目录名保持一致
	不要跟标准库冲突

*/

import (
	"fmt"
)

func main() {
	var a_g_e int64 = 19
	fmt.Println(a_g_e)
}
