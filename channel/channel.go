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
			time.Sleep(time.Second * 3)
			fmt.Println("n", n)
			if n > 3 {
				close(chnaO)
			} else {
				chnaO <- n
			}
		}
	}()

	go func() {
		n := 0
		for {
			n += 1
			time.Sleep(time.Second * 4)
			fmt.Println("n", n)
			if n > 3 {
				close(chnaO)
			} else {
				chnaO <- n
			}
		}
	}()

	go func() {
		for k := range chnaO {
			fmt.Println("1接收到了K", k)
		}
		fmt.Println("1进程结束")
	}()

	go func() {
		for k := range chnaO {
			fmt.Println("2接收到了K", k)
		}
		fmt.Println("2进程结束")
	}()

	time.Sleep(time.Second * 5000)
}
