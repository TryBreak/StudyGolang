package main // package 进行包的声明，建议：包的声明这个包和所在的文件夹同名
// 2. main 包是程序的入口包  一般 main 函数会放在这个包下面

import (
	fmt "fmt"

	utils "test.com/dbutils"

	test "test.com/test"
)

func init() {
	fmt.Println("main 函数的 init ")
}

func main() {
	var age int64 = 64
	fmt.Println("这里是main", age)
	fmt.Println("这里是test下的变量", test.StuNo) // 首字母必须大写才能被访问
	utils.Add()
	utils.Bbb()

	// 同级别下的源文件包的声明必须一致

	/*
		1. 使用包的原因：
			a. 我们不可能把所有的函数放在同一个原文件中，可以分门别类的把不同的函数放在不同的原文件中
			b. 解决同名问题：两个人都想定义同一个同名函数，字啊同一个文件中是不可以定义相同名字函数的。此时可以用包来区分
		2. 包的开始路径从 go.mod 开始的

		细节部分：
		1. package 声明包，建议：包的声明和这个包所在文件夹同名
		2. main包是程的入口包，一般main函数会放在这个包下
		3. 打包语法：
			package 包名
		4. 引入包的语法： import “包的路径”
		5. 有多个包，建议一次性导入格式如下：
		import(
				"fmt"
				"xxxx/xxxx/xxxx/xxxx"
		)
		6. 在函数调用的时候前面要定位到所在的包
		7. 首字母大写，函数可以被其它包访问
		8. 一个目录下不能有重复的函数
		9. 包名和文件夹的名字，可以不一样
		10. 一个目录下的同级文件归属一个包
		11. 包到底是什么：
			a. 在程序层面，所有使用相同 package 包名 的源文件组成的代码块
			c. 在源文件层面就是一个文件夹

		12. 包可以起别名，用了别名之后原来的包名就不能再用了

	*/
}
