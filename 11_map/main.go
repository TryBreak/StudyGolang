package main

import "fmt"

func main() {
	var m1 map[string]int
	fmt.Println(m1 == nil) // nil 没有在内存中开辟空间

	m1 = make(map[string]int, 10) // 要估算好该map容量,避免在程序运行期间再动态扩容
	m1["理想"] = 18                 //设置
	m1["mark"] = 27
	m1["姬无命"] = 35
	fmt.Println(m1)

	fmt.Println(m1["娜扎"])  //拿到了0值
	fmt.Println(m1["姬无命"]) //拿到了35值

	//约定成俗此处用ok接收
	v, ok := m1["姬无命"]
	if !ok {
		fmt.Println("查无此人")
	} else {
		fmt.Println(v)
	}
	fmt.Println("-----------")

	//遍历
	for k, v := range m1 {
		fmt.Printf("k:%v v:%v \n", k, v)
	}

	//删除
	delete(m1, "姬无命")
	fmt.Printf("删除: %v \n", m1)
	delete(m1, "娜扎")
	fmt.Printf("删除娜扎: %v \n", m1)

}
