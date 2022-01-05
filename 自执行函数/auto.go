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

	for i := 2; i > 0; i-- {
		log.Println("返回错误 xxx ，连接失败，正在重连...", i)
		time.Sleep(time.Second)
	}
	start()
}
