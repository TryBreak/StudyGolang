package main

import "fmt"

func demo1() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("aar[2:6] = ", arr[2:6])
	fmt.Println("aar[:6] = ", arr[:6])
	fmt.Println("aar[2:] = ", arr[2:])
	fmt.Println("aar[:] = ", arr[:])
}

func printSlice(s []int) {
	fmt.Printf("v:%v , len:%v , cap:%v\n", s, len(s), cap(s))
}

func curd() {
	//创建
	s := make([]int, 16, 128)
	s1 := []int{2, 4, 6, 8, 10}
	printSlice(s)

	fmt.Println("copy")
	copy(s,s1)
	printSlice(s)

	fmt.Println("新增")
	s = append(s, 10)
	printSlice(s)

	fmt.Println("删除")
	s = append(s[:3], s[4:]...)
	printSlice(s)

}

func main() {
	// demo1()
	curd()
}
