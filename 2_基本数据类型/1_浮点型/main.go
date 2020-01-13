package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Printf("%f\n", math.MaxFloat32) // go支持的最大的32位浮点数
	fmt.Printf("%f\n", math.Pi)
	fmt.Printf("%.2f\n", math.Pi) // 保留两位
}
