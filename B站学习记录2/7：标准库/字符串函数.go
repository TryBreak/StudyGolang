package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
字符串函数
*/

func main() {

	fmt.Println("------len函数使用-----")

	var str string = "golang你好"

	fmt.Println("str 字符串长度 -----", len(str))

	fmt.Println("------对字符串进行遍历-----")

	for i, value := range str {
		fmt.Printf("索引为--%d，值为--%c \n", i, value)
	}

	fmt.Println("------切片-----")

	var r []rune = []rune(str)
	fmt.Println("r -----", len(r), r)
	for i := 0; i < len(r); i++ {
		fmt.Printf("索引为--%d，值为--%c \n", i, r[i])
	}

	fmt.Println("------字符串转整数-----")
	var n, err = strconv.Atoi(str)
	fmt.Println("n---", n)
	fmt.Println("err---", err)

	fmt.Println("------整数转字符串-----")
	var s1 string = strconv.Itoa(999)
	fmt.Println("s1---", s1+"你好")

	fmt.Println("------统计一个字符串有几个指定的子串-----")
	var count int = strings.Count("golag and java ga", "g")
	fmt.Println("count---", count)

	fmt.Println("------查找子串是否在指定的字符串中-----")
	var isFind bool = strings.Contains("golag and java ga", "i")
	fmt.Println("isFind---", isFind)

	fmt.Println("------不区分大小写的字符串比较-----")
	var isTong bool = strings.EqualFold("GO", "go")
	fmt.Println("isTong---", isTong)

	fmt.Println("------字符串第一次出现的位置-----")
	var indeOf int = strings.Index("golag and java ga", "a")
	fmt.Println("indeOf---", indeOf)

	fmt.Println("------字符串的替换-----")
	var repla string = strings.Replace("golang and java gogogo", "go", "golang", -1)
	fmt.Println("repla---", repla)

	fmt.Println("------字符串切割标识符-----")
	var spli []string = strings.Split("go-python-java", "-")
	fmt.Println("spli---", spli)

	fmt.Println("------大小写字符串-----")
	var lower string = strings.ToLower("go-python-java")
	var upper string = strings.ToUpper("go-python-java")
	fmt.Println("lower---", lower)
	fmt.Println("upper---", upper)

	fmt.Println("------去掉两边指定字符串-----")
	var s4 string = strings.Trim("~go-python-java~", "~")
	fmt.Println("s4---", s4)

	var s5 string = strings.TrimLeft("~go-python-java~", "~")
	fmt.Println("s5---", s5)

	var s6 string = strings.TrimRight("~go-python-java~", "~")
	fmt.Println("s6---", s6)

	// 是否制定开头，是否指定结尾

}
