package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	timeStamp := 1648219200

	t := time.Unix(int64(timeStamp), 0)

	date := t.Format("15:04")

	numDate := strings.Replace(date, ":", ".", 1)

	fmt.Println(numDate)

}
