/*
  慕课网爬虫

*/

package main

import (
	"fmt"

	// "interface.com/infra"
	"interface.com/testing"
)

func getRetriever() testing.Retriever {
	return testing.Retriever{}
}

// ? : Something that can "Get"
type retriever interface {
	Get(string) string
}

/*

type -- 定义自定义类型

retriever -- 自定义类型的名字

interface -- 接口
		方法的集合

struct -- 结构体
		字段的集合

*/

func main() {
	var r retriever = getRetriever()
	// retriever := getRetriever()
	fmt.Println(r.Get("https://www.imooc.com"))
}
