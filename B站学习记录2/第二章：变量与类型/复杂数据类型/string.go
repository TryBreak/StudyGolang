package main

import "fmt"

func main() {

	var s1 string = "你好 go 语言"
	fmt.Println(s1)

	// 字符串是不可变的，字符串一单定义好，其中的字符值不能改变
	var s2 string = "abc"
	s2 = "def"
	// s2[0] = 't'   // 这样做会报错
	fmt.Println(s2)

	var s3 string = "`"
	fmt.Println(s3)

	// 字符串拼接
	var s5 string = "asd-" +
		"-123" +
		"-123" +
		"-123" +
		"-123" +
		"-123"
	fmt.Println(s5)
}
