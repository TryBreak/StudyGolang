package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	tttt()

	fmt.Println("22222")
	time.Sleep(9999999)
	fmt.Println("1111")
}

func tttt() {
	t := time.NewTicker(time.Second * 2)

	go func() {
		for range t.C {
			// 逻辑
			log.Println("发送ping")
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 1)
			// 这里是返回数据
			t.Reset(time.Second * 2)
		}
	}()
}
