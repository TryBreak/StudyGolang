package main

func main() {

	/*
		一共有26张图片 , 两个一组进行展示
		其中有4张是Yoda大师,
		每隔10分钟切换一次 ,
		每次切换都是完全随机 , 但同一组中不可能会出现同一张图片
		请问最少要多少次才可以切换成一组中两张都是Yoda?
	*/

	// 构建壁纸集合 0 代表Yoda
	bzhi := []int{}
	for i := 1; i <= 26; i++ {
		if i >= 23 {
			bzhi = append(bzhi, 0)
		} else {
			bzhi = append(bzhi, i)
		}
	}
	//进行递归循环,直到

}
