package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func demo() {
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

func traversal() {
	// 构建学生集合
	rand.Seed(time.Now().UnixNano())

	var scoreMap = make(map[string]int, 200)
	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("学生%02d号", i) //生成 `学生` 开头字符串
		value := rand.Intn(100)          // 生成 0 - 100 随机整数
		scoreMap[key] = value
	}
	// fmt.Println(scoreMap)
	//按照指定顺序遍历

	//取出map中的所有Key存入切片
	var keys = make([]string, 0, 200)
	for key := range scoreMap {
		keys = append(keys, key)
	}
	// fmt.Println(keys)

	//对切片进行排序
	sort.Strings(keys)
	// fmt.Println(keys)

	//按照排序后的key遍历map
	for _, key := range keys {
		fmt.Println(key, scoreMap[key])
	}
}

/*

%d就是普通的输出了
%2d是将数字按宽度为2，采用右对齐方式输出，若数据位数不到2位，则左边补空格
%02d，和%2d差不多，只不过左边补0
%.2d没见过，但从执行效果来看，和%02d一样

*/
func main() {
	// demo();
	traversal()
}
