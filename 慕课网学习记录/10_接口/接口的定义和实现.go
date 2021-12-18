package main

import "fmt"

// 接口由使用者定义

type Retiever interface {
	Get(url string) string
}

func download(r Retiever) string {
	return r.Get("www.imooc.com")
}

func main() {

	var r Retiever

	download(r)
	fmt.Println("墨七")
}
