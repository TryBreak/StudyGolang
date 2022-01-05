package main

import "fmt"

func main() {
	// param := ` {
	// 	"event": "error",
	// 	"msg": "Invalid sign",
	// 	"code": "60007"
	// }`

	c := []byte("")

	fmt.Println(len(c))
	fmt.Println("---------------------")

	str := `{"event":"login",`
	fmt.Println(len([]byte(str)))
}
