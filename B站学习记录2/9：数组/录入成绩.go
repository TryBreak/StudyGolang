package main

import "fmt"

func main() {

	var scores [5]int

	for i := 0; i < len(scores); i++ {
		fmt.Printf("请输入第%d学生的成绩 \n", i)
		fmt.Scanln(&scores[i])
	}
	fmt.Printf("所有学生的成绩是: %v \n", scores)

	fmt.Println("--------range 便利-----------------")
	// 数组,切片,字符串,map以及通道
	for key, val := range scores {
		fmt.Printf("key的值为 %v,val的值为%v \n", key, val)
	}

}
