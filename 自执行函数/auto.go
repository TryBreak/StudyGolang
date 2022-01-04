package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	start()
}

func start() {
	fmt.Println("123")

	n := 2

	for {
		log.Println("返回错误 xxx ，连接失败，正在重连...", n)
		n--
		time.Sleep(time.Second)
		if n == 0 {
			log.Println("重新连接中", n)
			start()
		}
	}
}
