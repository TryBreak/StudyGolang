package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Yes我爱慕课网!"
	for i, b := range []byte(s) { // b is byte
		fmt.Printf("(%v,%v) ", i, b) //16进制数
	}
	fmt.Println()

	for i, ch := range s { // ch is rune   // 这里 i 不是一一对应关系
		fmt.Printf("(%v,%v) ", i, ch)
	}
	fmt.Println()

	fmt.Println(
		"rune count:",
		utf8.RuneCountInString(s),
	)

	bytes := []byte(s)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%v ", ch) //单个字符输出
	}
	fmt.Println()

	for i, ch := range []rune(s) { //这里转化成rune 类型之后 下标就可以和汉字一一对应
		fmt.Printf("(%d,%c) ", i, ch)
	}

	fmt.Println()
	runeS2 := []rune(s)

	fmt.Println(string(runeS2[4]))

}
