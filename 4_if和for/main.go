package main

import (
	"fmt"
	"strconv"
)

func ifDemo1() {
	score := 65
	if score >= 90 {
		fmt.Println("A")
	} else if score > 75 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}
}

func ifDemo2() {
	if score := 690; score >= 90 {
		fmt.Println("A")
	} else if score > 75 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}
}

func forDemo99() {
	var all string
	for a := 1; a < 10; a++ {
		for b := 1; b < 10; b++ {
			vStr := strconv.Itoa(a * b)
			aStr := strconv.Itoa(a)
			bStr := strconv.Itoa(b)
			str := aStr + "*" + bStr + "=" + vStr + " "
			if a == b {
				str += "\n"
			}
			if a >= b {
				all += str
			}
		}
	}
	fmt.Println(all)
}

func forDemo() {
	//完整循环
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	// 省略初始语句
	i := 0
	for ; i < 10; i++ {
		fmt.Println(i)
	}
	// 省略初始和结束语句
	b := 0
	for b < 10 {
		fmt.Println(i)
		b++
	}
	//无限循环
	for {
		//循环体语句
	}
}

func forRange() {
	// Go语言中可以使用 `for range` 遍历数组、切片、字符串、map 及通道 （channel）
	s := []string{"a", "好", "c", "a我", "e", "f"}
	for index, item := range s {
		fmt.Printf("%v --- %v \n", index, item)
	}

	for _, item := range s {
		fmt.Printf("%v  \n", item)
	}
}

func switchDemo1() {
	week := 5
	str := ""
	switch week {
	case 1:
		str = "星期一"
	case 2:
		str = "星期二"
	case 3:
		str = "星期三"
	case 4:
		str = "星期四"
	case 5:
		str = "星期五"
	case 6:
		str = "星期六"
	case 7:
		str = "星期日"
	default:
		str = "err"
	}
	fmt.Println(str)
}

func switchDemo2() {
	//一个分支可以有多个值，多个case值中间使用英文逗号分隔。
	switch n := 7; n {
	case 1, 3, 5, 7, 9:
		fmt.Println("奇数")
	case 2, 4, 6, 8:
		fmt.Println("偶数")
	default:
		fmt.Println(n)
	}
}

func switchDemo3() {
	age := 29
	switch {
	case age < 25:
		fmt.Println("好好学习吧")
	case age > 25 && age < 30:
		fmt.Println("好好工作吧")
	case age > 60:
		fmt.Println("好好享受吧")
	default:
		fmt.Println("活着真好")

	}
}

func switchDemo4() {
	s := "b"
	switch {
	case s == "a":
		fmt.Println("a")
	case s == "b":
		fmt.Println("b")
		fallthrough // 执行符合条件的下一个case
	case s == "c":
		fmt.Println("c")
	default:
		fmt.Println("...")
	}
}

func gotoDemo() {
	// 直接终止当前的玄幻跳往某个标签
	for index := 0; index < 10; index++ {
		if index == 5 {
			goto breakTag
		}
		fmt.Println(index)
	}

breakTag:
	fmt.Println("结束for循环")

}

func for99Down() {
	for a := 9; a > 0; a-- {
		for b := 9; b > 0; b-- {
			if a >= b {
				fmt.Printf("%v*%v=%v ", a, b, a*b)
			}
		}
		fmt.Println()
	}
}

func main() {
	// ifDemo1()
	// ifDemo2()
	forDemo99()
	// forDemo()
	// forRange()
	// switchDemo1()
	// switchDemo2()
	// switchDemo3()
	// switchDemo4()
	// gotoDemo()
	for99Down()

}
