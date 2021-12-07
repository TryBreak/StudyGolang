package test

import "fmt"

var StuNo int64 = 20034 // 同一个包里面，不能定义重名函数或者变量

func init() {
	fmt.Println("test 函数的 init ")
}
