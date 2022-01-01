package main

import (
	"log"
	"time"
)

// 延迟函数

/*
	for {
		log.Println(time.Duration(1) * time.Millisecond * 1000 * 2)
		time.After(time.Duration(1) * time.Millisecond * 1000)
		time.Sleep(time.Duration(1) * time.Millisecond * 1000 * 2)
	}
*/

// 协程
// 轻量级线程

func main() {
	var a [10]int
	for i := 0; i < 10; i++ {
		go func(i int) {
			a[i] += i
		}(i)
		time.Sleep(time.Duration(1) * time.Millisecond * 1000 * 1)
		log.Println(a)
	}
	log.Println(a)
}
