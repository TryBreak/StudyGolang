package main

import "fmt"

func arrDemo() {
	//类型和容量是数组类型的一部分
	var a1 [3]bool // [true false true]
	var a2 [4]bool // [true false true true]
	fmt.Printf("a1:%T  a2%T \n", a1, a2)

	/*
		s数组的定义"
		var 数组变量名 [元素数量]T


		数组的初始化:
		如果不初始化: 默认都是零值(
			布尔值就是false,
			整形和浮点型就是 0 ,
			字符串就是 ""
		)
	*/
}

func arrInit() {
	// 初始化方式1
	var a1 = [3]bool{true, true, true}
	fmt.Println("a1", a1)

	//初始化方式2 : 根据推断数组长度
	a10 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("a10", a10)

	//初始化方式3 : 默认补零
	a3 := [5]int{1, 2}
	fmt.Println("a3", a3) // a3 [1 2 0 0 0]

	//初始化方式4 : 根据索引初始化
	a4 := [5]int{0: 1, 4: 2}
	fmt.Println("a4", a4) // a4 [1 0 0 0 2]
}

func arrFor() {
	citys := [...]string{"北京", "上海", "深圳"}
	//1. 根据索引遍历 : 和 javascript 一样
	for i := 0; i < len(citys); i++ {
		fmt.Println(citys[i])
	}
	//2. for range
	for i, v := range citys {
		fmt.Println(i, v)
	}
}

func arrMany() {
	// 多维数组
	// [[1 2] [3 4] [5 6]]
	var a1 [3][2]int
	/*
		[3] : 长度
		[2]int : 值的类型
	*/

	a1 = [3][2]int{
		[2]int{1, 2},
		[2]int{3, 4},
		[2]int{5, 6},
	}

	// fmt.Println(a1)  // [[1 2] [3 4] [5 6]]

	//多维数组遍历
	for _, v := range a1 {
		fmt.Println(v)
		for _, v2 := range v {
			fmt.Println(v2)
		}
	}

}

func arrType() {
	//值类型 : 数组copy 不会改变原数组
	a1 := [3]int{1, 2, 3}
	a2 := a1
	a2[1] = 100

	fmt.Println("a1", a1) // a1 [1 2 3]
	fmt.Println("a2", a2) // a2 [1 100 3]

}

func topic1() {
	//求数组[1, 3, 5, 7, 8]所有元素的和
	a := [...]int{1, 3, 5, 7, 8}
	num := 0
	for _, v := range a {
		num += v
	}
	fmt.Println(num)
}

func topic2() {
	/*
		找出数组中和为指定值的两个元素的下标，
		比如从数组[1, 3, 5, 7, 8]中找出和为8的两个元素的下标分别为(0,3)和(1,2)。
	*/

	a1 := [...]int{1, 3, 5, 7, 8}

	/*
		定义两个for
		第一个从第一个开始遍历
		内层从外层后面哪个开始
	*/

	for i := range a1 {
		for j := i + 1; j < len(a1); j++ {
			if a1[i]+a1[j] == 8 {
				fmt.Printf("(%v , %v) \n", i, j)
			}
		}
	}
}

func main() {
	// arrDemo()
	// arrInit()
	// arrFor()
	// arrMany()
	// arrType()
	// topic1()
	topic2()
}
