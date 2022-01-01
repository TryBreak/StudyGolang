package main

import (
	"log"
	"time"
)

func channelDemo() {
	c := make(chan int)

	go func() {
		for {
			n := <-c
			log.Println(n)
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 1)
			c <- 1
			log.Println(c)
			time.Sleep(time.Second * 1)
			c <- 2
			log.Println(c)

		}
	}()
	//  2022/01/02 02:35:39
	// wg.Wait()
}

func main() {
	channelDemo()
	time.Sleep(time.Second * 20)
}
