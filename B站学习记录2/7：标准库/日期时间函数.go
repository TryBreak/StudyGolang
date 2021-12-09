/*
日期跟时间函数
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("----日期跟时间函数------")

	var now time.Time = time.Now()

	fmt.Printf("%v ~~ 对应的类型是：%T \n", now, now)
	fmt.Printf("年：%v \n", now.Year())
	fmt.Printf("月：%v \n", int(now.Month()))
	fmt.Printf("日：%v \n", now.Day())
	fmt.Printf("时：%v \n", now.Hour())
	fmt.Printf("分：%v \n", now.Minute())
	fmt.Printf("秒：%v \n", now.Second())

	fmt.Println("----指定格式------")
	t := now.Format("2006-01-02 15:04:05")

	fmt.Printf("t 的值为：%v \n", t)
	// 20060102150405
}
