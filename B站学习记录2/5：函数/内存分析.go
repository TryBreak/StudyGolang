/*
内存的逻辑区分

栈 - 基本数据类型

		main 函数栈帧
		为函数再次开辟一个栈帧
		函数执行完毕，则销毁栈帧释放内存

堆 - 复杂数据类型

代码区- 代码本身

*/

package main

import "fmt"

func main() {
	fmt.Println("内存分析")
}
