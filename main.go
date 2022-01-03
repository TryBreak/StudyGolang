package main

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
type Animal interface {
	Speak() string
}

type Dog struct{}

func (d Dog) Speak() string {
	return "Woof!"
}

func main() {
}
