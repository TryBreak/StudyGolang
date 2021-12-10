package main

import "fmt"

func main() {

	fmt.Println("--------第一种---------")
	var arr1 [3]int = [3]int{3, 6, 9}
	fmt.Printf("arr1 的值为:%v \n", arr1)

	fmt.Println("--------第二种---------")
	var arr2 = [3]int{1, 4, 7}
	fmt.Printf("arr2 的值为:%v \n", arr2)

	fmt.Println("--------第三种---------")
	var arr3 = [...]int{1, 2, 5, 7, 8, 6}
	fmt.Printf("arr3 的值为:%v \n", arr3)

	fmt.Println("--------第四种---------")
	var arr4 = [...]int{2: 66, 0: 33, 3: 65, 5: 87, 8, 6}
	fmt.Printf("arr4 的值为:%v \n", arr4)
}
