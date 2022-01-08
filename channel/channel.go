package main

import (
	"fmt"
	"time"
)

func main() {
	chnaO := make(chan int)

	go func() {
		n := 0
		for {
			n += 1
			if n > 6 {
				close(chnaO)
				break
			}
			fmt.Printf("通道发送")
			chnaO <- n
			time.Sleep(time.Second * 1)
		}
	}()
	// https://cloud.tencent.com/developer/article/1412490
	go func() {
		for k := range chnaO {
			if k == 2 {
				chnaO = nil
			}
			fmt.Println("K:", k)
		}
		fmt.Println("1进程结束")
	}()

	go func() {
		for k := range chnaO {
			fmt.Println("接收到了K", k)
		}
		fmt.Println("K进程结束")
	}()

	go func() {
		for {
			select {
			case x := <-chnaO:
				fmt.Println("select接收到了 x ", x)
			}
		}
	}()

	time.Sleep(time.Second * 5000)
}
