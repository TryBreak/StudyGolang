package main

// 导入语句

/**
import "fmt"
*/
import (
	"fmt"
	// "time"
)

//函数外只能放置标识符(变量\常量\函数\类型)的声明
//	fmt.Printf("hello, world") //非法,会报错

// 程序入口
func main() {
	fmt.Printf("hello, world")
	// time.Sleep(5 * time.Second)  命令行延迟五秒关闭
}
