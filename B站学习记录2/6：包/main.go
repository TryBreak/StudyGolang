package main // package 进行包的声明，建议：包的声明这个包和所在的文件夹同名
// 2. main 包是程序的入口包  一般 main 函数会放在这个包下面

import (
	"fmt"

	test "test.com/test"
)

func main() {
	var age int64 = 64
	fmt.Println("这里是main", age)
	fmt.Println("这里是test下的变量", test.StuNo)
}
