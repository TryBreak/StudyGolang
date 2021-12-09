package main

import "fmt"

func main() {
	fmt.Println("------数组的引入------")
	// 五个学生的成绩求出平均数

	// 定义数组
	var scores [5]int
	// 数组的赋值
	scores[0] = 95
	scores[1] = 91
	scores[2] = 39
	scores[3] = 60
	scores[4] = 21

	// 求和
	var sum int = 0
	for i := 0; i < len(scores); i++ {
		sum += scores[i]
	}

	// 平均数
	var avg int = sum / len(scores)

	// 输出
	fmt.Printf("成绩总和为：%v，成绩平均数为：%v \n", sum, avg)

}
