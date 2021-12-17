package main

import "fmt"

func main() {
	test()
	fmt.Println("上面的触发操作执行成功")
	fmt.Println("正常执行下面的逻辑")

	fmt.Println("--defer--recover--")

}

func test() {
	// 调用 defer + recover 来捕获错误
	defer func() { //压入堆栈中，最后处理
		// 调用 recover 内置函数，可以捕获错误
		err := recover() // 捕获错误
		// 如果没有捕获错误，返回值为零值： nil
		if err != nil { // err == nil 表示没有错误，如果不等于 nil 就表示有错误
			fmt.Println("错误已经捕获")
			fmt.Println("err是：", err)
		}
	}()

	num1 := 100
	num2 := 0
	result := num1 / num2
	fmt.Println(result)

}
