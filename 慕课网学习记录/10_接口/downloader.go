/*
  慕课网爬虫

*/

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func retrieve(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	bytes, _ := ioutil.ReadAll(resp.Body)

	return bytes
}

func main() {
	bytes := retrieve("https://www.imooc.com")
	fmt.Printf("%s \n", bytes)
}
