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

func main() {
	var r retriever = getRetriever()
	// retriever := getRetriever()
	fmt.Println(r.Get("https://www.imooc.com"))
}
