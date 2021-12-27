package main

import mypackage "myproject/mypackage"

func main() {
	// 调用add方法
	mypackage.Add(10, 25)
	// 调用 books结构体
	var threeKindom mypackage.Book
	threeKindom.Title = "threekindom"
	threeKindom.Author = "andy"
	// threekindom.bookid = 1003 这里不能定义，会报错，小写开头只供包内调用
}
