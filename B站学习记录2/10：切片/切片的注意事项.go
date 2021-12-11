package main

import "fmt"

/*
1. 切片不能直接使用，需要让其引用到一个数组或者make一个空间共切片使用
2. 切片不能越界
//不能超过数组的长度

3. 简写方式

var slice = arr[0:end]  ===> arr[:end]


4. 切片可以继续切片


5. 切片可以动态增长

*/

func main() {
	fmt.Println("-------切片可以动态增长----------")

	var arrint [6]int = [6]int{1, 4, 7, 2, 5, 8}

	var slice []int = arrint[1:4]

	slice2 := append(slice, 88, 50) //进行扩容，老数组扩容为新数组
	// 创建一个新数组，将老数组的值放到新数组中，在数组中追加88，50
	fmt.Println(len(slice), slice)
	fmt.Println(len(slice2), slice2)

	fmt.Printf("slice 数组的地址：%p \n", &slice)
	fmt.Printf("slice2 数组的地址：%p \n", &slice2)
	slice = append(slice, 88, 50) //内存地址不变
	fmt.Printf("slice 扩容数组的地址：%p \n", &slice)

	slice = append(slice, slice2...) // 合并数组v

	fmt.Printf("slice 扩容数组的地址：%p \n", &slice)
	fmt.Printf("slice ：%v \n", slice)

	fmt.Println("-------拷贝----------")

	var a []int = []int{1, 4, 7, 3, 6, 9}
	var b []int = make([]int, 10)

	// 拷贝
	copy(b, a) // 将A的内容复制到B中
	fmt.Println(a)
	fmt.Println(b)

}
