package main

import "fmt"

func swap(a, b *int) {
	*a, *b = *b, *a
}

func main() {
	a, b := 3, 4
	swap(&a, &b)
	fmt.Printf("a:%v , b%v", a, b)

	/// & 取地址    *根据地址取值
}
