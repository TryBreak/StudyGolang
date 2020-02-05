package main

import "fmt"

func demo1() {
	m := map[string]string{
		"name":    "没常量",
		"course":  "golang",
		"site":    "imooc",
		"quality": "notbad",
	}
	m2 := make(map[string]int, 10) // m2 == empty map 内存中开辟了空间

	var m3 map[string]int // m3 == nil  没有在内存中开辟空间
	fmt.Println(m, m2, m3)
	fmt.Println("Traversing map")
	for k, v := range m {
		fmt.Println(k, v)
	}
	fmt.Println("Getting values")
	course, ok := m["course"]
	fmt.Println(course, ok)
	couse, ok := m["couse"]
	fmt.Println(couse, ok)

	fmt.Println("----Deleting values----")
	name, ok := m["name"]
	fmt.Println(name, ok)

	delete(m, "name")
	name, ok = m["name"]
	fmt.Println(name, ok)
	fmt.Println(len(m))

}

func main() {
	demo1()
}
