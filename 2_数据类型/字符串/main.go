package main

import "fmt"

func main() {
	/*
		双引号为字符串
		单引号为字符
		其它的和一般的程序没差别
		字符串拼接 使用加号,如:  str1+str2
		len(str) // 长度
	*/
	// 	var str = `
	// 111
	// 222
	// 333
	// 44
	// 	`
	// fmt.Print(str)

	/*
		基础库  string

		string.Split(str,"/") //分割

		string.Contains(str,"xx") // 包含

		除此之外基础库还包含:
		前缀后缀判断  // 以xxx开头,以xxx结尾
		子串出现的位置 // 某个字符串在另一个字符串里面的位置
		合并操作
		切片操作 : str[2]


	*/

	/*
			字符类型:
		byte   => 一般的英文字符
		rune  => 中文字符


	*/
	//字符串修改
	s2 := "白萝卜"
	s3 := []rune(s2)
	s3[0] = '红' //此处为单引号
	fmt.Println(string(s3))
}
