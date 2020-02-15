package main

import (
	"fmt"

	api "helloworld/api" //helloworld  是项目名 go.mod 中声明的

	alse "helloworld/alse"

	"github.com/labstack/echo"
)

func main() {
	ret := alse.Add(5, 8)
	fmt.Println(ret)

	ret2 := alse.Add(8, 12)
	fmt.Println("apis1", ret2)

	ret3 := alse.Add(8, 13)
	fmt.Println("apis2", ret3)

	ret4 := alse.Add(8, 18)
	fmt.Println("apis3", ret4)

	e := echo.New()
	e.GET("/", api.HelloWorld)
	e.Logger.Fatal(e.Start(":1122"))
}
