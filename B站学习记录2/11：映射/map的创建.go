package main

import "fmt"

func main() {
	fmt.Println("-----方式一--------")
	var a map[int]string

	a = make(map[int]string, 10)

	a[12344] = "张三"
	a[45] = "李四"
	a['c'] = "王麻子"
	a['d'] = "无主物"
	a['e'] = "王柳"

	fmt.Println(a)

	fmt.Println("-----方式二--------")
	b := make(map[int]string)
	b['d'] = "无主物"
	b['e'] = "王柳"
	fmt.Println(b)

	fmt.Println("-----方式三--------")
	c := map[string]string{
		"a":  "张三",
		"b":  "李四",
		"cs": "王五",
		"我":  "赵六",
	}

	fmt.Println(c)

}
