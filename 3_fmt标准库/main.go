package main

import "fmt"

func main() {
	const n = 100
	fmt.Printf("%T\n", n) // 打印类型
	fmt.Printf("%v\n", n) // 打印值
	fmt.Printf("%b\n", n) // 二进制
	fmt.Printf("%o\n", n) // 八进制
	fmt.Printf("%d\n", n) // 十进制
	fmt.Printf("%x\n", n) // 十六进制

	const s = "梅花"
	fmt.Printf("%s\n", s)  // 字符串
	fmt.Printf("%#v\n", s) // # 添加描述符号
	fmt.Printf("%p \n", s) // p 打印内存地址
}

/*

%d就是普通的输出了
%2d是将数字按宽度为2，采用右对齐方式输出，若数据位数不到2位，则左边补空格
%02d，和%2d差不多，只不过左边补0
%.2d没见过，但从执行效果来看，和%02d一样

*/
