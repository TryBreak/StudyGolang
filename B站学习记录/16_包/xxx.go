package main

import (
	"fmt"

	calc "./calc/calc.go"
)

func main() {
	ret := calc.Add(10, 30)
	fmt.Println(ret)
}
