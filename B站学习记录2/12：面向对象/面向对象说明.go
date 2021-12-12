/*

【1】golang 语言面向对象编程说明：
1. golang 也支持面向对象变编程（OOP），但是和传统的面向对象有区别，
并不是纯粹的面向对象语言。所以我们说Golang支持面向对象的特性是比较准确的。

2. Golang 没有类（class），Go语言的结构体（struct）和其它编程语言的类（class）
有同等的地位，你可以理解golang是基于 struct 来实现OOP特性的。

3. golang面向对象编程非常简洁，去掉了传统的OOP语言的方法重载、构造函数和析构函数、
隐藏的this指针等等

4. golang 仍然有面向对象编程的继承，封装和多态的特性，只是实现方式和其他语言OOP
不一样，比如继承golang没有 extends 关键字，继承是通过匿名字段来实现。


【2】结构体的引入



*/

package main

import "fmt"

func main() {

	/*
		具体对象：
		一位老师：
			姓名：珊珊
			年龄：31
			性别：女

	*/

	fmt.Println("-----变量表示--------")

	// 赵姗姗老师
	var name string = "赵姗姗"
	var age int = 31
	var sex string = "女"

	//  马士兵老师
	var name1 string = "赵姗姗"
	var age2 int = 31
	var sex3 string = "女"

	// name4 ....
	// 不利于数据的管理与维护
	// 很多属性属于一个对象，用变量管理太分散了
	// 数组切片 都不合适

	/*
		结构体：

		属性：字段

		动作/行为/方法
	*/

}
