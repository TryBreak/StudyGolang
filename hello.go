package main

// 导入语句

/**
import "fmt"
*/
import (
	"fmt"
	"time"
	// "time"
)

//函数外只能放置标识符(变量\常量\函数\类型)的声明
//	fmt.Printf("hello, world") //非法,会报错

// 程序入口
func main() {
	// fmt.Printf("hello, world \n")
	// time.Sleep(5 * time.Second)  命令行延迟五秒关闭
	start()
}

func start() {
	num := 9000000000
	startTime := time.Now().UnixNano() / 1e6

	// for i := 0; i < num; i++ {
	// 	// fmt.Printf("当前i的值:%v \n", i)
	// }

	for i := num; i >= 0; i-- {
		// fmt.Printf("当前i的值:%v \n", i)
	}

	endTime := time.Now().UnixNano() / 1e6

	fmt.Printf("当的时间差:%v \n", endTime-startTime)

}
