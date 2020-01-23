package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	/*
		一共有26张图片 , 两个一组进行展示
		其中有4张是Yoda大师,
		每隔10分钟切换一次 ,
		每次切换都是完全随机 , 但同一组中不可能会出现同一张图片
		请问最少要多少次才可以切换成一组中两张都是Yoda?
	*/

	// 构建壁纸集合 0 代表Yoda

	for i := 1; i <= 26; i++ {
		if i >= 23 {
			bzhi = append(bzhi, 0)
		} else {
			bzhi = append(bzhi, i)
		}
	}
	/*
		1. 生成一组随机下标
		2. 如果下标相等则重新生成直到不等
		3. 随即范围是 0-25
	*/
	/*
		根据随机数遍历图片数组,都为0则停下 , 记录一次
	*/
	for index := 0; index < 100; index++ {
		start()
	}
	//求平均数
	allNum := 0
	for index := 0; index < len(law); index++ {
		allNum += law[index]
	}
	fmt.Printf("则百次匹配Yoda,平均用了%v次", allNum/len(law))

}

var (
	bzhi []int
	law  []int
)

func start() {
	num := 0
	for {
		a := rand.Intn(25)
		b := rand.Intn(25)
		if a != b {
			num++
			bzhiA := bzhi[a]
			bzhiB := bzhi[b]
			// fmt.Printf("第一次匹配,壁: %v ; 壁纸: %v \n", bzhiA, bzhiB)
			if bzhiA == bzhiB {
				// fmt.Printf("成功匹配到两个Yoda , 一共用了%v次 \n", num)
				law = append(law, num)
				break
			}
		}
	}
}
