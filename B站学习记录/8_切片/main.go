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

func sliceDemo1() {
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
	var a1 = [...]int{1, 2, 3, 4, 5, 6, 7, 89}
	a3 := a1[0:4] // 基于一个数组进行切割  从 下标0到下标4 (左闭右开) 左包含右不包含
	a4 := a1[1:6]
	a5 := a4[1:6]
	a6 := a4[3:7] // 右边没写就是最后一个
	a7 := a4[:4]  // 左边没写就是最开头开始
	a8 := a4[:]   // [0:len(4)]
	fmt.Println("a3", a3)
	fmt.Println("a4", a4)
	fmt.Println("a5", a5)
	fmt.Println("a6", a6)
	fmt.Println("a7", a7)
	fmt.Println("a8", a8)

	//切片的容量(cap)指的是底层数组的容量 , 从第一个元素到最后一个元素的数量
	fmt.Printf("len(a5):%v cap(a5):%v \n", len(a5), cap(a5))
	fmt.Printf("len(a4):%v cap(a4):%v \n", len(a4), cap(a4))
	//切片属于引用类型,全部指向底层数组
}
func sliceDemo2() {
	//make()  函数创造切片
	// 长度为5 容量为10
	s1 := make([]int, 5, 10)
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d \n", s1, len(s1), cap(s1))

	// 底层数组长度为10
	s2 := make([]int, 0, 10)
	fmt.Printf("s2=%v len(s2)=%d cap(s2)=%d \n", s2, len(s2), cap(s2))

}

func sliceCopy() {
	var a1 = [...]int{1, 2, 3, 4, 5, 6, 7, 9, 10}
	fmt.Printf("a1=%v len(a1)=%d cap(a1)=%d \n", a1, len(a1), cap(a1))

	s3 := []int{1, 2, 3, 4, 5}
	s4 := s3
	fmt.Println("s4,s3", s4, s3)
	s3[1] = 1000
	fmt.Println("s4,s3", s4, s3)
	s5 := s3[1:3]
	fmt.Println("s4,s3,s5", s4, s3, s5)
	s4[2] = 999
	fmt.Println("s4,s3,s5", s4, s3, s5)
	s5[0] = 777
	fmt.Println("s4,s3,s5", s4, s3, s5)

	a2 := a1[1:]
	fmt.Println("a1,a2", a1, a2)
	a2[0] = 789
	fmt.Println("a1,a2", a1, a2)

}

func appendDemo() {
	// 切片扩容
	s1 := []string{"北京", "上海", "深圳"}

	fmt.Printf("s1=%v len(s1)=%v cap(s1)=%v \n", s1, len(s1), cap(s1))

	s1[2] = "广州"

	fmt.Printf("s1=%v len(s1)=%v cap(s1)=%v \n", s1, len(s1), cap(s1))

	//必须使用变量去接受 , 然后替换原数组
	s1 = append(s1, "北海道")
	fmt.Printf("s1=%v len(s1)=%v cap(s1)=%v \n", s1, len(s1), cap(s1))
	s1 = append(s1, "成都", "杭州")
	fmt.Printf("s1=%v len(s1)=%v cap(s1)=%v \n", s1, len(s1), cap(s1))
	s1 = append(s1, "武汉")
	fmt.Printf("s1=%v len(s1)=%v cap(s1)=%v \n", s1, len(s1), cap(s1))
	ssa := []string{"哈哈", "西安", "樱花街道"}
	fmt.Printf("ssa:%v \n", ssa)
	s1 = append(s1, ssa...) // ...表示拆开
	fmt.Printf("s1=%v len(s1)=%v cap(s1)=%v \n", s1, len(s1), cap(s1))
	/*
		数组内存放不下,所以就扩容了
		cap = 6
	*/

	//基本上只能用于切片而不能用于数组
	ss := [5]int{1, 2, 3, 4, 5}
	ss1 := ss[1:]
	ss1 = append(ss1, 7)

	fmt.Printf("ss1=%v len(ss1)=%v cap(ss1)=%v \n", ss1, len(ss1), cap(ss1))
	fmt.Printf("ss=%v len(ss)=%v cap(ss)=%v \n", ss, len(ss), cap(ss))
	fmt.Printf("ss:%T ss1:%T \n", ss, ss1)
}

func copyDemo() {
	a1 := []int{1, 3}
	a2 := a1
	var a3 = make([]int, 3, 3)
	copy(a3, a1)
	fmt.Println(a1, a2, a3)
	a1[0] = 100 //a3 不变
	fmt.Println(a1, a2, a3)
	copy(a3, a1)
	fmt.Println(a1, a2, a3)

}

func delDemo() {
	// 从切片中删除元素
	a := []int{1, 2, 3, 4, 5, 6, 7, 8}
	// 要删除索引为2的元素
	a = append(a[:2], a[3:]...) // 删除元素会改变底层数组
	fmt.Println(a)
}

func delArrDemo() {
	a1 := [...]int{1, 2, 3, 4, 5}
	s1 := a1[:]

	/*
		1. 切片不保存具体的值
		2. 切片对应一个底层数组
		3. 底层数组都是占用一块连续的内容

		所以 , 切片删除 , 则底层数组对应移位  ,而连续内存不变
	*/

	s1 = append(s1[:2], s1[3:]...) // 修改了底层数组

	s1[0] = 100 //修改底层数组

	fmt.Printf("a1=%v len(a1)=%v cap(a1)=%v \n", a1, len(a1), cap(a1))
	fmt.Printf("s1=%v len(s1)=%v cap(s1)=%v \n", s1, len(s1), cap(s1))
	fmt.Printf("s1%p \n", s1)
}

func main() {
	// sliceDemo1()
	// sliceDemo2()
	// sliceCopy()
	// appendDemo()
	// copyDemo()
	// delDemo()
	delArrDemo()
}
