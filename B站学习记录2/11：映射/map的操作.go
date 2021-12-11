package main

import "fmt"

func main() {

	var bMap = make(map[int]string)
	fmt.Println("-----增-------")
	bMap[12] = "张三"
	bMap[34] = "王五12"
	bMap[33] = "王五2"
	bMap[32] = "王五3"

	fmt.Println("-----删-------")

	delete(bMap, 34)
	fmt.Println(bMap)

	fmt.Println("-----改-------")

	bMap[33] = "赵六"
	fmt.Println(bMap)

	fmt.Println("-----查-------")
	v, flag := bMap[5654]

	fmt.Println(v, flag)

	fmt.Println("-----获取长度-------")
	fmt.Println(len(bMap))

	fmt.Println("-----遍历:只支持 for-range-------")

	for k, v := range bMap {
		fmt.Println(k, v)
	}

}
