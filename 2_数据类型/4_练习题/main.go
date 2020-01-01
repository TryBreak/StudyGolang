package main

import (
	"fmt"
	"unicode/utf8"
)

/*

1. 编写代码分别定义一个整型、浮点型、布尔型、字符串型变量，使用fmt.Printf()搭配%T分别打印出上述变量的值和类型。
2. 编写代码统计出字符串 "hello我的人生理想" 中汉字的数量。

*/

func main() {
	// printT()
	chartNum()
}

func printT() {
	const (
		a = 100
		b = 3.1415926
		s = "我是一个大帅比"
	)
	var c bool

	fmt.Printf("a: %T -- %v \n", a, a) // 整型 int
	fmt.Printf("b: %T -- %v \n", b, b) // 浮点型 float64
	fmt.Printf("c: %T -- %v \n", c, c) // 布尔  bool
	fmt.Printf("s: %T -- %v \n", s, s) // 字符串 string

}

func chartNum() {
	//长度计算
	const s = "hello我的人生理想"
	fmt.Printf("s: %T -- %v \n", s, s)
	fmt.Println(len(s)) // 23

	s1 := []byte(s)

	fmt.Println(len(s1)) // 23

	s2 := []rune(s)

	fmt.Println(len(s2)) // 11

	fmt.Println(utf8.RuneCountInString(s)) // 11

	fmt.Println("--------------") // 分隔符

	for i := 0; i < len(s); i++ {
		fmt.Printf("%v ( %c \n", s[i], s[i])
	}
	fmt.Println("--------------") // 分隔符

	var num = 0
	for _, r := range s {
		// fmt.Printf("%v ( %c \n", r, r)
		var intStr = int(r)
		if intStr > 128 {
			num++
		}
		// fmt.Println(intStr)
	}
	fmt.Printf("该字符串中一共有%v 个中文字符", num)

}
