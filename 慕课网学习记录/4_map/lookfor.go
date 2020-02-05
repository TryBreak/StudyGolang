package main

import "fmt"

// 例子: 寻找最长不含有重复字符串的子串

/*
	abcabcbb  ->  abc
	bbbbbbb  ->  b
	pwwkew  -> wke
*/

func lengthOfNunRepeatingSubStr(s string) int {
	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0
	for i, ch := range []rune(s) { // 这里利用的map key的属性 , 不断地记录每个字符串最后一个的位置
		lastI, ok := lastOccurred[ch]
		if ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}

func main() {
	s := "一二三二一"
	lent := lengthOfNunRepeatingSubStr(s)
	fmt.Println(lent)
}
