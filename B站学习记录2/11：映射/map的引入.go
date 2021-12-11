/*

映射，map，Go语言内置的一种类型，
它的键值相关联，我们可以通过键 Key 来获取对应的值 value。
类似其它语言的集合

基本语法
var map[keytype]valuetype

PS:key，value的类型：bool、数字、string、指针、channel、
还可以只是包含前面几个类型的接口、结构体、数组\

PS：key通常为 int、string类型，value通常为数字（整数、浮点数）
string、map、结构体

PS：slice、map、function 不可以


*/

package main

import "fmt"

func main() {
	fmt.Println("----------------------")

	var a map[int]string
	// 只声明 map，是没有分配内存空间的
	fmt.Printf("a 的内存地址:%p ,值: %v,类型:%T \n",
		&a, a, a)
	fmt.Printf("a 的长度:%v  \n", len(a))

	// 必须通过 make 函数进行初始化,才会分配空间:

	fmt.Println("--------make--------------")

	a = make(map[int]string, 10)
	a[20095452] = "张三"
	a[20095387] = "李四"
	a[200957291] = "王五"
	a[200912342] = "赵六"
	//map 是无序的键值对
	fmt.Printf("a 的内存地址:%p ,值: %v,类型:%T \n",
		&a, a, a)
	fmt.Printf("a 的长度:%v  \n", len(a))

}

/*
特点:
1. map 集合使用前一定要 make
2. key-value无序
3. key 不能重复,后面的会替换前面的

*/
