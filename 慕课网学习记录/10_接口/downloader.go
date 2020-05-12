/*
  慕课网爬虫

*/

package main

import (
	"fmt"

	"interface.com/infra"
)

func getRetriever() infra.Retriever {
	return infra.Retriever{}
}

func main() {
	var retriever infra.Retriever = getRetriever()
	// retriever := getRetriever()
	fmt.Println(retriever.Get("https://www.imooc.com"))
}
