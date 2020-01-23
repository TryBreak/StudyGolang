package main

import (
	"fmt"
	"strconv"
)

func main() {
	fun1() // 启动递归

}

var a = 1 //声明递归控制
var b = 1
var all = ""

func fun1() {

	b = 1
	fun2()
	// fmt.Println(a) // 打印a查看变量变化情况

	a++         //控制变量自增 ， 以达到条件值
	if a < 10 { // 判断是否符合自调用条件
		fun1() //自调用
	} else {
		fmt.Println(all) //查看保存的结果
	}
}

// fun2 的控制变量
func fun2() {
	// fmt.Println(b) // 打印b查看变量变化情况
	// fmt.Printf("a:%v -- b%v\n", a, b)  // 检测 a  b 碰面情况
	var aStr = strconv.Itoa(a)
	var bStr = strconv.Itoa(b)
	var vStr = strconv.Itoa(a * b)

	var str = aStr + "*" + bStr + "=" + vStr + " " // 合成式子

	if b == a { //换行的条件
		str = str + "\n" // 添加换行
	}

	if b > a {
		//不要了
	} else {
		//要
		all += str // 保存式子
	}

	// fmt.Println(str) // 查看式子是否正确

	b++
	if b < 10 {
		fun2() // 自调用
	}
}
