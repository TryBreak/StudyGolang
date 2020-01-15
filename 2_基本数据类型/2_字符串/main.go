package main

import "fmt"

func main() {
	// demo()
	splitStr()
}

func splitStr() {
	s := "[ 9 +0]=  9"

	var s1 string

	for _, v := range s {
		// fmt.Printf("%s \n", string(v))
		if string(v) == " " {
			// fmt.Printf("这玩意儿是空格\n")
		} else {
			s1 = s1 + string(v)
		}
	}

	fmt.Println(s1)
}

func demo() {
	/*
		双引号为字符串
		单引号为字符
		其它的和一般的程序没差别

		字符串拼接 使用加号,如:  str1+str2

	*/

	//多行字符串
	var str = `
	111
	222
	333
	44

		`

	fmt.Print(str)

	/*

		字符串方法
		len(str) // 长度

		基础库  string
		string.Split(str,"/") //分割
		string.Contains(str,"xx") // 包含
		除此之外基础库还包含:

		strings.HasPrefix,
		strings.HasSuffix    // 以xxx开头,以xxx结尾
		strings.Index(),strings.LastIndex()   // 子串出现的位置 // 某个字符串在另一个字符串里面的位置
		strings.Join(a[]string, sep string) 合并操作
		切片操作 : str[2]

	*/

	/*
			字符类型:
		byte   => 一般的英文字符
		rune  => 中文字符
	*/
	//字符串修改
	s2 := "白萝卜\n"
	s3 := []rune(s2)
	s3[0] = '红' //此处为单引号
	fmt.Printf(string(s3))

	fmt.Printf("xxxxxxx\n")

	for _, r := range s3 {
		fmt.Printf("%s \n", string(r))
	}

	changeString()

}

//修改字符串
func changeString() {
	s1 := "bing"
	//强制类型转换
	byteS1 := []byte(s1)
	byteS1[0] = 'p'
	fmt.Println(string(byteS1))

	s2 := "白萝卜"
	runeS2 := []rune(s2)
	runeS2[0] = '红'
	fmt.Println(string(runeS2))
}
