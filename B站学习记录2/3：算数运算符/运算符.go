package main

import "fmt"

func main() {

	// + 号
	// 1.正数  2. 相加操作  3.字符串拼接
	var n1 int = +10
	var n2 int = 4 + 7
	var s1 string = "abc" + "def"

	fmt.Println(n1)
	fmt.Println(n2)
	fmt.Println(s1)

	fmt.Println("--------------")

	// 除号
	fmt.Println(10 / 3)    // int 类型运算，一定是整数类型
	fmt.Println(10.00 / 3) // 浮点参与运算，一定是浮点类型
	fmt.Println("--------------")

	//  % 取模  等价公式  a%b= a-a/b*b

	fmt.Println(10 % 3) // 10-10/3*3  10 -9
	fmt.Println(-10 % 3)
	fmt.Println(10 % -3)
	fmt.Println(-10 % -3)
	fmt.Println("--------------")

	// ++ 自增  -- 自减
	var a1 int64 = 100
	a1++
	a1++
	fmt.Println(a1)
	a1--
	a1--
	a1--
	a1--
	fmt.Println(a1)
	fmt.Println("--------------")
	// --a   ++a  是违法的

	var num3 int = 10
	num3 += 20 // num3 = num3 +20
	fmt.Println(num3)
	fmt.Println("--------------")

	// 交换两个数值并输出结果
	var a int = 8
	var b int = 4
	fmt.Printf("a = %v,b = %v \n", a, b)

	var c int
	c = a
	a = b
	b = c
	fmt.Printf("a = %v,b = %v \n", a, b)
	fmt.Println("--------------")

	// 关系运算符
	fmt.Println(5 == 9)
	fmt.Println(5 != 9)
	fmt.Println(5 > 9)
	fmt.Println(5 < 9)
	fmt.Println(5 >= 9)
	fmt.Println(5 <= 9)
	fmt.Println("--------------")
	// 逻辑运算符
	//  && 与   短路与 --- 只要第一个数值结果是 false 后面表达式就不参与运算
	//  || 或		短路或 --- 只要有一个是 true 后面的就不用计算了
	// ！ 非

	fmt.Println(true && 1 > 2)
	fmt.Println(false && true)
	fmt.Println(true && false)

	// 其它运算符
	// &  返回变量的地址
	// *  返回指针变量的数值

}
