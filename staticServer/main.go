package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	httpserver()
}

func httpserver() {

	//文件浏览
	fmt.Println("start")
	http.Handle("/", http.FileServer(http.Dir("www")))
	fmt.Println("访问 localhost:3007")
	err := http.ListenAndServe(":3007", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
