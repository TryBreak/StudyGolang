package api

import (
	"net/http"

	"github.com/labstack/echo"

	"fmt"

	alse "helloworld/alse"
)

//HelloWorld is a func
func HelloWorld(c echo.Context) error {
	ret := alse.Add(8, 10)

	fmt.Println("apis truexxxx", ret)
	fmt.Println("c", c)

	return c.JSON(http.StatusOK, "hello world")
}
