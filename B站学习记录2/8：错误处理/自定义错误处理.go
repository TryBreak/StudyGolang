package main

import (
	"errors"
	"fmt"
)

// 自定义错误要调用 errors 包

func main() {
	fmt.Println("--自定义错误捕获--")

	err := test()
	if err != nil { // == nil 则表示没错
		fmt.Println("自定义错误", err)
		panic(err) // 后面的程序不再执行
	}
	fmt.Println("上面的触发操作执行成功")
	fmt.Println("正常执行下面的逻辑")

}

func test() (err error) {
	num1 := 10
	num2 := 0
	if num2 == 0 {
		// 抛出错误
		return errors.New("除数不能为0哦~~")
	} else {
		// 除数不为零则打印结果
		result := num1 / num2
		fmt.Println(result)
		return nil
	}
}
