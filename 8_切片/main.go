package main

import "fmt"

/*
	数组有很多局限性 , 所以有了切片(Slice)
	切片（Slice）是一个拥有相同类型元素的可变长度的序列。
	它是基于数组类型做的一层封装。它非常灵活，支持自动扩容。
	它的内部结构包含 `地址`、`长度`和`容量`

	var name []T

	name:表示变量名
	T:表示切片中的元素类型
*/
func main() {
	//跟数组的区别就是没有定义长度
	var s1 []int    //	int类型切片
	var s2 []string //  string类型切片
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Printf("%T \n", s1)
	fmt.Println(s1 == nil) //true     nil 的意思是空 和 nil 的概念相似
	fmt.Println(s2 == nil) //true
	s1 = []int{1, 2, 3}
	s2 = []string{"北京", "广州", "上海", "深圳"}
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s1 == nil) //false
	fmt.Println(s2 == nil) //false
	//切片的长度 len()
	fmt.Printf("len(s1):%v  cpa(s1)%v \n", len(s1), cap(s1)) //3   3
	fmt.Printf("len(s2):%v  cpa(s2)%v \n", len(s2), cap(s2)) //4   4

	//数组转化为切片
	var a1 = [...]int{1, 2, 3, 4, 5, 6, 6, 7, 89}
	a3 := a1[0:4] // 基于一个数组进行切割  从 下标0到下标4 (左闭右开) 左包含右不包含
	a4 := a1[1:6]
	a5 := a4[1:6]
	fmt.Println(a3)
	fmt.Println(a4)
	fmt.Println(a5)

}
